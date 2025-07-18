package modelg

import (
	"context"
	"database/sql"
	"iter"
)

type DB interface {
	Exec(ctx context.Context, query string, args ...any) (Result, error)
	Query(ctx context.Context, query string, args ...any) iter.Seq2[Row, error]
	CreateVariablesScope() QueryVariablesScope
}

type TxDB interface {
	DB
	Begin(ctx context.Context, opts *TxOpts) (Tx, error)
}

type Tx interface {
	Exec(ctx context.Context, query string, args ...any) (Result, error)
	Query(ctx context.Context, query string, args ...any) iter.Seq2[Row, error]
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type TxOpts struct {
	ReadOnly  bool
	Isolation sql.IsolationLevel
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Columns() ([]string, error)
	Scan(dest ...any) error
}
