package modelg

import (
	"context"
	"database/sql"
	"iter"
)

type SQLDB interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type simpleDB struct {
	db SQLDB
}

var _ DB = (*simpleDB)(nil)

func NewSimpleDB(db SQLDB) DB {
	return &simpleDB{db: db}
}

func (s *simpleDB) Exec(ctx context.Context, query string, args ...any) (Result, error) {
	return s.db.ExecContext(ctx, query, args...)
}

func (s *simpleDB) Query(ctx context.Context, query string, args ...any) iter.Seq2[Row, error] {
	return func(yield func(Row, error) bool) {
		rows, err := s.db.QueryContext(ctx, query, args...)
		if err != nil {
			yield(nil, err)
			return
		}
		closed := false
		defer func() {
			if !closed {
				rows.Close()
			}
		}()

		for rows.Next() {
			if !yield(sqlRow{row: rows}, nil) {
				return
			}
		}

		if err := rows.Err(); err != nil {
			yield(nil, err)
			return
		}

		closed = true
		if err := rows.Close(); err != nil {
			yield(nil, err)
			return
		}
	}
}

func (db *simpleDB) CreateVariablesScope() QueryVariablesScope {
	return NewGenericQueryVariablesScope()
}

type sqlRow struct {
	row *sql.Rows
}

func (r sqlRow) Columns() ([]string, error) {
	return r.row.Columns()
}

func (r sqlRow) Scan(dest ...any) error {
	return r.row.Scan(dest...)
}
