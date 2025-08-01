package modelg

import (
	"context"
	"iter"
)

// A UniqueConstraintError is an error that indicates a unique constraint violation in the database.
// Depending on the database, information about the violated constraint may be available in the ConstraintColumns fields.
type UniqueConstraintError struct {
	ConstraintColumns []*ColumnInfo
	underlying        error
}

func NewUniqueConstraintError(err error, constraintColumns []*ColumnInfo) *UniqueConstraintError {
	return &UniqueConstraintError{
		ConstraintColumns: constraintColumns,
		underlying:        err,
	}
}

func (e *UniqueConstraintError) Error() string {
	return e.underlying.Error()
}

func (e *UniqueConstraintError) Unwrap() error {
	return e.underlying
}

// A CheckConstraintError is an error that indicates a check constraint violation in the database.
type CheckConstraintError struct {
	ConstraintName string
	underlying     error
}

func NewCheckConstraintError(err error, constraintName string) *CheckConstraintError {
	return &CheckConstraintError{
		ConstraintName: constraintName,
		underlying:     err,
	}
}

func (e *CheckConstraintError) Error() string {
	return e.underlying.Error()
}

func (e *CheckConstraintError) Unwrap() error {
	return e.underlying
}

// A NotNullConstraintError is an error that indicates a NOT NULL constraint violation in the database.
type NotNullConstraintError struct {
	ConstraintColumn *ColumnInfo
	underlying       error
}

func NewNotNullConstraintError(err error, constraintColumn *ColumnInfo) *NotNullConstraintError {
	return &NotNullConstraintError{
		ConstraintColumn: constraintColumn,
		underlying:       err,
	}
}

func (e *NotNullConstraintError) Error() string {
	return e.underlying.Error()
}

func (e *NotNullConstraintError) Unwrap() error {
	return e.underlying
}

// A ForeignKeyConstraintError is an error that indicates a foreign key constraint violation in the database.
type ForeignKeyConstraintError struct {
	ParentTable   string
	ChildTable    string
	ModifiedTable string // The table that was modified in the operation that caused the error.
	underlying    error
}

func NewForeignKeyConstraintError(err error, childTableName, parentTableName, modifiedTableName string) *ForeignKeyConstraintError {
	return &ForeignKeyConstraintError{
		ParentTable:   parentTableName,
		ChildTable:    childTableName,
		ModifiedTable: modifiedTableName,
		underlying:    err,
	}
}

func (e *ForeignKeyConstraintError) Error() string {
	return e.underlying.Error()
}

func (e *ForeignKeyConstraintError) Unwrap() error {
	return e.underlying
}

type ColumnInfo struct {
	Name   string
	Table  string
	Schema string
}

type ErrorTranslator func(error) error

type ErrorTranslatingDB struct {
	DB
	translateError ErrorTranslator
}

func NewErrorTranslatingDB(underlyingDB DB, translateError ErrorTranslator) *ErrorTranslatingDB {
	return &ErrorTranslatingDB{
		DB:             underlyingDB,
		translateError: translateError,
	}
}

func (db *ErrorTranslatingDB) Unwrap(target any) bool {
	return unwrapOrDelegrate(db.DB, target)
}

func (db *ErrorTranslatingDB) Exec(ctx context.Context, query string, args ...any) (Result, error) {
	result, err := db.DB.Exec(ctx, query, args...)
	return result, db.translateError(err)
}

func (db *ErrorTranslatingDB) Query(ctx context.Context, query string, args ...any) iter.Seq2[Row, error] {
	return func(yield func(Row, error) bool) {
		rowIter := db.DB.Query(ctx, query, args...)
		for row, err := range rowIter {
			if err != nil {
				yield(nil, db.translateError(err))
				return
			}
			if !yield(row, nil) {
				return
			}
		}
	}
}

type ErrorTranslatingTxDB struct {
	txdb TxDB
	*ErrorTranslatingDB
}

func NewErrorTranslatingTxDB(underlyingDB TxDB, translateError ErrorTranslator) *ErrorTranslatingTxDB {
	return &ErrorTranslatingTxDB{
		txdb:               underlyingDB,
		ErrorTranslatingDB: NewErrorTranslatingDB(underlyingDB, translateError),
	}
}

func (db *ErrorTranslatingTxDB) Begin(ctx context.Context, opts *TxOpts) (Tx, error) {
	tx, err := db.txdb.Begin(ctx, opts)
	if err != nil {
		return nil, db.translateError(err)
	}
	return &ErrorTranslatingTx{
		tx:             tx,
		translateError: db.translateError,
	}, nil
}

type ErrorTranslatingTx struct {
	tx             Tx
	translateError ErrorTranslator
}

func (db *ErrorTranslatingTx) Unwrap(target any) bool {
	return unwrapOrDelegrate(db.tx, target)
}

func (tx *ErrorTranslatingTx) Exec(ctx context.Context, query string, args ...any) (Result, error) {
	result, err := tx.tx.Exec(ctx, query, args...)
	return result, tx.translateError(err)
}

func (tx *ErrorTranslatingTx) Query(ctx context.Context, query string, args ...any) iter.Seq2[Row, error] {
	return func(yield func(Row, error) bool) {
		rowIter := tx.tx.Query(ctx, query, args...)
		for row, err := range rowIter {
			if err != nil {
				yield(nil, tx.translateError(err))
				return
			}
			if !yield(row, nil) {
				return
			}
		}
	}
}

func (tx *ErrorTranslatingTx) Commit(ctx context.Context) error {
	if err := tx.tx.Commit(ctx); err != nil {
		return tx.translateError(err)
	}
	return nil
}

func (tx *ErrorTranslatingTx) Rollback(ctx context.Context) error {
	if err := tx.tx.Rollback(ctx); err != nil {
		return tx.translateError(err)
	}
	return nil
}
