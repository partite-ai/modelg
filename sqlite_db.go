package modelg

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"iter"
	"regexp"
	"strings"
	"time"

	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

type SqlitePool interface {
	Take(ctx context.Context) (*sqlite.Conn, error)
	Put(conn *sqlite.Conn)
}

type SqliteDB struct {
	pool SqlitePool
}

var _ TxDB = (*SqliteDB)(nil)

func NewSqliteDB(pool SqlitePool) *SqliteDB {
	return &SqliteDB{
		pool: pool,
	}
}

func (db *SqliteDB) Raw() any {
	return db.pool
}

func (db *SqliteDB) Exec(ctx context.Context, query string, args ...any) (Result, error) {
	conn, err := db.pool.Take(ctx)
	if err != nil {
		return nil, err
	}
	defer db.pool.Put(conn)

	return sqliteExec(ctx, db, conn, query, args...)
}

func (db *SqliteDB) Query(ctx context.Context, query string, args ...any) iter.Seq2[Row, error] {
	return sqliteQuery(ctx, db, func() (*sqlite.Conn, func(), error) {
		conn, err := db.pool.Take(ctx)
		if err != nil {
			return nil, nil, err
		}
		return conn, func() { db.pool.Put(conn) }, nil
	}, query, args...)
}

func (db *SqliteDB) Begin(ctx context.Context, txOpts *TxOpts) (Tx, error) {
	readonly := txOpts != nil && txOpts.ReadOnly

	var startCmd string
	if readonly {
		startCmd = "BEGIN DEFERRED;"
	} else {
		startCmd = "BEGIN IMMEDIATE;"
	}

	conn, err := db.pool.Take(ctx)
	if err != nil {
		return nil, err
	}

	inited := false
	defer func() {
		if !inited {
			db.pool.Put(conn)
		}
	}()

	if readonly {
		err = sqlitex.Execute(conn, "PRAGMA query_only = ON;", nil)
		if err != nil {
			return nil, fmt.Errorf("failed to set read-only mode: %w", err)
		}
	} else {
		err = sqlitex.Execute(conn, "PRAGMA query_only = OFF;", nil)
		if err != nil {
			return nil, fmt.Errorf("failed to set read-write mode: %w", err)
		}
	}

	err = sqlitex.Execute(conn, startCmd, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	inited = true
	return &sqliteTx{
		conn: conn,
		db:   db,
	}, nil
}

func (db *SqliteDB) CreateVariablesScope() QueryVariablesScope {
	return NewGenericQueryVariablesScope()
}

func (db *SqliteDB) prepareStatement(conn *sqlite.Conn, query string, args ...any) (*sqlite.Stmt, error) {
	stmt, err := conn.Prepare(strings.TrimSpace(query))
	if err != nil {
		return nil, err
	}

	needDestroy := false
	defer func() {
		if needDestroy {
			stmt.Finalize()
		}
	}()

	bindParams := stmt.BindParamCount()
	if len(args) != bindParams {
		return nil, fmt.Errorf("expected %d arguments, got %d", bindParams, len(args))
	}

	for i, arg := range args {
		val, err := driver.DefaultParameterConverter.ConvertValue(arg)
		if err != nil {
			return nil, fmt.Errorf("failed to convert argument %d: %w", i, err)
		}

		switch v := val.(type) {
		case int64:
			stmt.BindInt64(i+1, v)
		case float64:
			stmt.BindFloat(i+1, v)
		case string:
			stmt.BindText(i+1, v)
		case []byte:
			stmt.BindBytes(i+1, v)
		case bool:
			stmt.BindBool(i+1, v)
		case time.Time:
			stmt.BindText(i+1, v.Format("2006-01-02 15:04:05.999999999-07:00"))
		case nil:
			stmt.BindNull(i + 1)
		default:
			return nil, fmt.Errorf("unsupported argument type %T for argument %d", arg, i)
		}
	}

	needDestroy = false
	return stmt, nil
}

type sqliteTx struct {
	db   *SqliteDB
	conn *sqlite.Conn
}

func (tx *sqliteTx) Exec(ctx context.Context, query string, args ...any) (Result, error) {
	oldInterrupt := tx.conn.SetInterrupt(ctx.Done())
	defer func() {
		tx.conn.SetInterrupt(oldInterrupt)
	}()

	stmt, err := tx.db.prepareStatement(tx.conn, query, args...)
	if err != nil {
		return nil, err
	}
	needsReset := true
	defer func() {
		if needsReset {
			stmt.Reset()
		}
	}()

	for {
		hasRow, err := stmt.Step()
		if err != nil {
			return nil, tx.db.translateError(err)
		}
		if !hasRow {
			break
		}
	}

	if err := stmt.Reset(); err != nil {
		return nil, fmt.Errorf("failed to reset statement: %w", err)
	}
	needsReset = false

	changeCount := tx.conn.Changes()
	lastInsertID := tx.conn.LastInsertRowID()
	return sqliteResult{lastInsertID: lastInsertID, rowsAffected: int64(changeCount)}, nil
}

func (tx *sqliteTx) Query(ctx context.Context, query string, args ...any) iter.Seq2[Row, error] {
	return func(yield func(Row, error) bool) {
		oldInterrupt := tx.conn.SetInterrupt(ctx.Done())
		defer func() {
			tx.conn.SetInterrupt(oldInterrupt)
		}()

		stmt, err := tx.db.prepareStatement(tx.conn, query, args...)
		if err != nil {
			yield(nil, err)
			return
		}
		needsReset := true
		defer func() {
			if needsReset {
				stmt.Reset()
			}
		}()

		var cols []string
		for {
			rowReturned, err := stmt.Step()
			if err != nil {
				yield(nil, tx.db.translateError(err))
				return
			}

			if !rowReturned {
				break
			}

			if !yield(sqliteRow{stmt: stmt, colPtr: &cols}, nil) {
				return
			}
		}

		if err := stmt.Reset(); err != nil {
			yield(nil, tx.db.translateError(err))
			return
		}
		needsReset = false
	}
}

func (t *sqliteTx) Commit(ctx context.Context) error {
	err := sqlitex.Execute(t.conn, "COMMIT;", nil)
	if err != nil {
		return err
	}
	t.db.pool.Put(t.conn)
	t.conn = nil
	return nil
}

func (t *sqliteTx) Rollback(ctx context.Context) error {
	if t.conn == nil {
		return nil
	}

	defer func() {
		t.db.pool.Put(t.conn)
		t.conn = nil
	}()

	oldDoneCh := t.conn.SetInterrupt(nil)
	defer t.conn.SetInterrupt(oldDoneCh)

	err := sqlitex.Execute(t.conn, "ROLLBACK;", nil)
	if err != nil {
		return err
	}
	return nil
}

func sqliteExec(ctx context.Context, db *SqliteDB, conn *sqlite.Conn, query string, args ...any) (Result, error) {
	oldInterrupt := conn.SetInterrupt(ctx.Done())
	defer func() {
		conn.SetInterrupt(oldInterrupt)
	}()

	stmt, err := db.prepareStatement(conn, query, args...)
	if err != nil {
		return nil, err
	}
	needsReset := true
	defer func() {
		if needsReset {
			stmt.Reset()
		}
	}()

	for {
		hasRow, err := stmt.Step()
		if err != nil {
			return nil, db.translateError(err)
		}
		if !hasRow {
			break
		}
	}

	if err := stmt.Reset(); err != nil {
		return nil, fmt.Errorf("failed to reset statement: %w", err)
	}
	needsReset = false

	changeCount := conn.Changes()
	lastInsertID := conn.LastInsertRowID()
	return sqliteResult{lastInsertID: lastInsertID, rowsAffected: int64(changeCount)}, nil
}

func sqliteQuery(ctx context.Context, db *SqliteDB, connf func() (*sqlite.Conn, func(), error), query string, args ...any) iter.Seq2[Row, error] {
	return func(yield func(Row, error) bool) {
		conn, release, err := connf()
		if err != nil {
			yield(nil, err)
			return
		}
		defer release()

		oldInterrupt := conn.SetInterrupt(ctx.Done())
		defer func() {
			conn.SetInterrupt(oldInterrupt)
		}()

		stmt, err := db.prepareStatement(conn, query, args...)
		if err != nil {
			yield(nil, err)
			return
		}
		needsReset := true
		defer func() {
			if needsReset {
				stmt.Reset()
			}
		}()

		var cols []string
		for {
			rowReturned, err := stmt.Step()
			if err != nil {
				yield(nil, db.translateError(err))
				return
			}

			if !rowReturned {
				break
			}

			if !yield(sqliteRow{stmt: stmt, colPtr: &cols}, nil) {
				return
			}
		}

		if err := stmt.Reset(); err != nil {
			yield(nil, db.translateError(err))
			return
		}
		needsReset = false
	}
}

type sqliteRow struct {
	stmt   *sqlite.Stmt
	colPtr *[]string
}

func (r sqliteRow) Columns() ([]string, error) {
	if *r.colPtr != nil {
		return *r.colPtr, nil
	}
	colCount := r.stmt.ColumnCount()
	cols := make([]string, colCount)
	for i := range colCount {
		colName := r.stmt.ColumnName(i)
		cols[i] = colName
	}
	*r.colPtr = cols
	return cols, nil
}

func (r sqliteRow) Scan(dest ...any) error {
	for i := 0; i < r.stmt.ColumnCount(); i++ {
		if i >= len(dest) {
			return fmt.Errorf("not enough destination arguments for %d columns", r.stmt.ColumnCount())
		}

		if r.stmt.ColumnIsNull(i) {
			err := ConvertAssign(dest[i], nil)
			if err != nil {
				return err
			}
			continue
		}

		colType := r.stmt.ColumnType(i)

		if _, ok := dest[i].(*time.Time); ok && colType == sqlite.TypeText {
			data := r.stmt.ColumnText(i)
			parsedTime, ok := parseSqliteTime(data)
			if !ok {
				return fmt.Errorf("failed to parse time from text column %d: %s", i, data)
			}
			if err := ConvertAssign(dest[i], parsedTime); err != nil {
				return fmt.Errorf("failed to convert text column %d: %w", i, err)
			}
			continue
		}

		switch colType {
		case sqlite.TypeBlob:
			colLen := r.stmt.ColumnLen(i)
			data := make([]byte, colLen)
			r.stmt.ColumnBytes(i, data)
			if err := ConvertAssign(dest[i], data); err != nil {
				return fmt.Errorf("failed to convert blob column %d: %w", i, err)
			}
		case sqlite.TypeText:
			data := r.stmt.ColumnText(i)
			if err := ConvertAssign(dest[i], data); err != nil {
				return fmt.Errorf("failed to convert text column %d: %w", i, err)
			}
		case sqlite.TypeInteger:
			data := r.stmt.ColumnInt64(i)
			if err := ConvertAssign(dest[i], data); err != nil {
				return fmt.Errorf("failed to convert integer column %d: %w", i, err)
			}
		case sqlite.TypeFloat:
			data := r.stmt.ColumnFloat(i)
			if err := ConvertAssign(dest[i], data); err != nil {
				return fmt.Errorf("failed to convert float column %d: %w", i, err)
			}
		}
	}
	return nil
}

type sqliteResult struct {
	lastInsertID int64
	rowsAffected int64
}

func (r sqliteResult) LastInsertId() (int64, error) {
	return r.lastInsertID, nil
}

func (r sqliteResult) RowsAffected() (int64, error) {
	return r.rowsAffected, nil
}

var parseTimeFormats = []string{
	"2006-01-02 15:04:05.999999999-07:00",
	"2006-01-02T15:04:05.999999999-07:00",
	"2006-01-02 15:04:05.999999999",
	"2006-01-02T15:04:05.999999999",
	"2006-01-02 15:04",
	"2006-01-02T15:04",
	"2006-01-02",
}

// Attempt to parse s as a time. Return (s, false) if s is not
// recognized as a valid time encoding.
func parseSqliteTime(s string) (any, bool) {
	ts := strings.TrimSuffix(s, "Z")

	for _, f := range parseTimeFormats {
		t, err := time.Parse(f, ts)
		if err == nil {
			return t, true
		}
	}

	return s, false
}

var sqliteConstraintErrorRegex = regexp.MustCompile(`constraint failed: ([^\:]*)$`)
var sqliteColumnsRegex = regexp.MustCompile(`([^,\s]+)`)

func (db *SqliteDB) translateError(err error) error {
	if err == nil {
		return nil
	}
	if resultCode := sqlite.ErrCode(err); resultCode != sqlite.ResultError {
		baseErr := errors.Unwrap(err)

		switch resultCode {
		case sqlite.ResultConstraintUnique:
			return NewUniqueConstraintError(err, sqlliteColumnInfoFromError(baseErr))
		case sqlite.ResultConstraintNotNull:
			columnInfo := sqlliteColumnInfoFromError(baseErr)
			if len(columnInfo) != 1 {
				return NewNotNullConstraintError(err, nil)
			}
			return NewNotNullConstraintError(err, columnInfo[0])
		case sqlite.ResultConstraintCheck:
			return NewCheckConstraintError(err, sqlliteConstraintNameFromError(baseErr))
		case sqlite.ResultConstraintForeignKey:
			return NewForeignKeyConstraintError(err, "", "", "")
		case sqlite.ResultConstraintTrigger:
			if strings.Contains(baseErr.Error(), "FOREIGN KEY constraint failed") {
				return NewForeignKeyConstraintError(err, "", "", "")
			}
		}
	}

	return err
}

func sqlliteColumnInfoFromError(err error) []*ColumnInfo {
	matches := sqliteConstraintErrorRegex.FindStringSubmatch(err.Error())

	if matches == nil {
		return nil
	}

	columnNames := matches[1]
	parts := sqliteColumnsRegex.FindAllString(columnNames, -1)

	var columnInfos []*ColumnInfo
	for _, col := range parts {
		colParts := strings.SplitN(col, ".", 2)
		if len(colParts) > 1 {
			columnInfos = append(columnInfos, &ColumnInfo{
				Name:  colParts[1],
				Table: colParts[0],
			})
		} else {
			columnInfos = append(columnInfos, &ColumnInfo{
				Name:  col,
				Table: "",
			})
		}
	}

	return columnInfos
}

func sqlliteConstraintNameFromError(err error) string {
	matches := sqliteConstraintErrorRegex.FindStringSubmatch(err.Error())

	if matches == nil {
		return ""
	}

	return matches[1]
}
