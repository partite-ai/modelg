package modelg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"iter"
	"regexp"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxPool interface {
	Acquire(ctx context.Context) (c *pgxpool.Conn, err error)
	BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)
}

type PgxDB struct {
	pool PgxPool
}

var _ TxDB = (*PgxDB)(nil)

func NewPgxDB(pool PgxPool) *PgxDB {
	return &PgxDB{
		pool: pool,
	}
}

func (db *PgxDB) Raw() any {
	return db.pool
}

func (db *PgxDB) Begin(ctx context.Context, opts *TxOpts) (Tx, error) {
	var pgxOpts pgx.TxOptions
	if opts != nil {
		if opts.ReadOnly {
			pgxOpts.AccessMode = pgx.ReadOnly
		} else {
			pgxOpts.AccessMode = pgx.ReadWrite
		}

		switch opts.Isolation {
		case sql.LevelReadUncommitted:
			pgxOpts.IsoLevel = pgx.ReadUncommitted
		case sql.LevelReadCommitted:
			pgxOpts.IsoLevel = pgx.ReadCommitted
		case sql.LevelWriteCommitted:
			return nil, fmt.Errorf("pgx does not support LevelWriteCommitted isolation level")
		case sql.LevelRepeatableRead:
			pgxOpts.IsoLevel = pgx.RepeatableRead
		case sql.LevelSnapshot:
			pgxOpts.IsoLevel = pgx.ReadCommitted
		case sql.LevelSerializable:
			pgxOpts.IsoLevel = pgx.Serializable
		case sql.LevelLinearizable:
			return nil, fmt.Errorf("pgx does not support LevelLinearizable isolation level")
		}
	}
	tx, err := db.pool.BeginTx(ctx, pgxOpts)
	if err != nil {
		return nil, err
	}
	return &pgxTx{
		db:   db,
		conn: tx,
	}, nil
}

func (db *PgxDB) Exec(ctx context.Context, query string, args ...any) (Result, error) {
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	return pgxExec(ctx, db, conn, query, args...)
}

func (db *PgxDB) Query(ctx context.Context, query string, args ...any) iter.Seq2[Row, error] {
	return pgxQuery(ctx, db, func() (pgxConn, func(), error) {
		conn, err := db.pool.Acquire(ctx)
		if err != nil {
			return nil, nil, err
		}
		return conn, conn.Release, nil
	}, query, args...)
}

func (db *PgxDB) CreateVariablesScope() QueryVariablesScope {
	return NewPostgresQueryVariablesScope()
}

type pgxTx struct {
	db   *PgxDB
	conn pgx.Tx
}

func (t *pgxTx) Exec(ctx context.Context, query string, args ...any) (Result, error) {
	return pgxExec(ctx, t.db, t.conn, query, args...)
}

func (t *pgxTx) Query(ctx context.Context, query string, args ...any) iter.Seq2[Row, error] {
	return pgxQuery(ctx, t.db, func() (pgxConn, func(), error) {
		return t.conn, func() {}, nil
	}, query, args...)
}

func (t *pgxTx) Commit(ctx context.Context) error {
	if err := t.conn.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (t *pgxTx) Rollback(ctx context.Context) error {
	if err := t.conn.Rollback(ctx); err != nil && err != pgx.ErrTxClosed {
		return fmt.Errorf("failed to rollback transaction: %w", err)
	}
	return nil
}

type pgxConn interface {
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

func pgxExec(ctx context.Context, db *PgxDB, conn pgxConn, query string, args ...any) (Result, error) {
	ct, err := conn.Exec(ctx, query, args...)
	if err != nil {
		return nil, db.translateError(ctx, err)
	}

	return pgxResult{ct: ct}, nil
}

func pgxQuery(ctx context.Context, db *PgxDB, connf func() (pgxConn, func(), error), query string, args ...any) iter.Seq2[Row, error] {
	return func(yield func(Row, error) bool) {
		conn, release, err := connf()
		if err != nil {
			yield(nil, err)
			return
		}
		defer release()
		rows, err := conn.Query(ctx, query, args...)
		if err != nil {
			yield(nil, db.translateError(ctx, err))
			return
		}
		defer rows.Close()

		var colNames []string
		for rows.Next() {
			if !yield(pgxRow{
				rows:        rows,
				ptrColNames: &colNames,
			}, nil) {
				return
			}
		}

		if err := rows.Err(); err != nil {
			yield(nil, db.translateError(ctx, err))
			return
		}
	}
}

type pgxRow struct {
	rows        pgx.Rows
	ptrColNames *[]string
}

func (r pgxRow) Columns() ([]string, error) {
	if *r.ptrColNames != nil {
		return *r.ptrColNames, nil
	}
	cols := r.rows.FieldDescriptions()
	columnNames := make([]string, len(cols))
	for i, col := range cols {
		columnNames[i] = string(col.Name)
	}
	*r.ptrColNames = columnNames
	return columnNames, nil
}

func (r pgxRow) Scan(dest ...any) error {
	if err := r.rows.Scan(dest...); err != nil {
		return err
	}
	return nil
}

type pgxResult struct {
	ct pgconn.CommandTag
}

func (r pgxResult) LastInsertId() (int64, error) {
	return 0, fmt.Errorf("pgx does not support LastInsertId")
}

func (r pgxResult) RowsAffected() (int64, error) {
	return r.ct.RowsAffected(), nil
}

func (db *PgxDB) translateError(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}
	if pgErr := (*pgconn.PgError)(nil); errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23001": // restrict_violation
			childTable, parentTable, modifiedTable := db.tryResolveForeignKeyInfo(ctx, pgErr)
			return NewForeignKeyConstraintError(err, childTable, parentTable, modifiedTable)
		case "23502": // not_null_violation
			var colInfo *ColumnInfo
			if pgErr.ColumnName != "" {
				colInfo = &ColumnInfo{
					Name:   pgErr.ColumnName,
					Table:  pgErr.TableName,
					Schema: pgErr.SchemaName,
				}
			}
			return NewNotNullConstraintError(err, colInfo)
		case "23503": // foreign_key_violation
			childTable, parentTable, modifiedTable := db.tryResolveForeignKeyInfo(ctx, pgErr)
			return NewForeignKeyConstraintError(err, childTable, parentTable, modifiedTable)
		case "23505": // unique_violation
			return NewUniqueConstraintError(err, db.tryResolveConstraintColumns(ctx, pgErr.ConstraintName))
		case "23514": // check_violation
			return NewCheckConstraintError(err, pgErr.ConstraintName)
		}
	}
	return err
}

var pgxFkInvalidChildRegex = regexp.MustCompile(`^insert or update on table "([^"]+)" violates foreign key constraint "([^"]+)"$`)
var pgxFkParentTableRegex = regexp.MustCompile(`^is not present in table "([^"]+)"$`)
var pgxFkRestrictRegex = regexp.MustCompile(`^update or delete on table "([^"]+)" violates RESTRICT setting of foreign key constraint "([^"]+)" on table "([^"]+)"$`)
var pgxFkInvalidParentRegex = regexp.MustCompile(`^update or delete on table "([^"]+)" violates foreign key constraint "([^"]+)" on table "([^"]+)"$`)

func (db *PgxDB) tryResolveForeignKeyInfo(ctx context.Context, pgerr *pgconn.PgError) (string, string, string) {
	if match := pgxFkInvalidChildRegex.FindStringSubmatch(pgerr.Message); match != nil {
		childTable := match[1]
		var parentTable string
		pgxFkParentTableRegex.FindStringSubmatch(pgerr.Detail)
		if match := pgxFkParentTableRegex.FindStringSubmatch(pgerr.Detail); match != nil {
			parentTable = match[1]
		}
		return childTable, parentTable, childTable
	} else if match := pgxFkRestrictRegex.FindStringSubmatch(pgerr.Message); match != nil {
		return match[3], match[1], match[1]
	} else if match := pgxFkInvalidParentRegex.FindStringSubmatch(pgerr.Message); match != nil {
		return match[3], match[1], match[1]
	}
	return "", "", ""
}

func (db *PgxDB) tryResolveConstraintColumns(ctx context.Context, name string) []*ColumnInfo {
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		return nil
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT table_schema, table_name, column_name  FROM information_schema.constraint_column_usage WHERE constraint_name = $1", name)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var columnInfos []*ColumnInfo
	for rows.Next() {
		var schemaName, tableName, columnName string
		if err := rows.Scan(&schemaName, &tableName, &columnName); err != nil {
			return nil
		}

		columnInfos = append(columnInfos, &ColumnInfo{
			Name:   columnName,
			Table:  tableName,
			Schema: schemaName,
		})
	}

	if len(columnInfos) == 0 {
		// Maybe the constraint is actually an index
		rows, err := conn.Query(ctx, `select
			ns.nspname,
			tc.relname,
			a.attname
		from pg_catalog.pg_class ic
		inner join pg_catalog.pg_index i on ic."oid" = i.indexrelid
		inner join pg_catalog.pg_class tc on i.indrelid = tc."oid"
		inner join pg_catalog.pg_attribute a on tc."oid" = a.attrelid and a.attnum = any(indkey[0:indnkeyatts-1])
		inner join pg_catalog.pg_namespace AS ns
			ON tc.relnamespace = ns.oid
		where ic.relname = $1 and ic.relkind = 'i'`, name)
		if err != nil {
			return nil
		}
		defer rows.Close()

		for rows.Next() {
			var schemaName, tableName, columnName string
			if err := rows.Scan(&schemaName, &tableName, &columnName); err != nil {
				return nil
			}

			if strings.HasPrefix(columnName, "$$") {
				// Skip hidden columns
				continue
			}

			columnInfos = append(columnInfos, &ColumnInfo{
				Name:   columnName,
				Table:  tableName,
				Schema: schemaName,
			})
		}
	}

	return columnInfos
}
