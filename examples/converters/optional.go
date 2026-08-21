package converters

import (
	"database/sql"
	"database/sql/driver"

	"github.com/partite-ai/modelg"
	"github.com/partite-ai/optional"
)

func OptionalScanner[T any](wrapped *optional.Optional[T], createScanner func(*T) sql.Scanner) *optionalScanner[T] {
	return &optionalScanner[T]{
		wrapped:       wrapped,
		createScanner: createScanner,
	}
}

type optionalScanner[T any] struct {
	wrapped       *optional.Optional[T]
	createScanner func(*T) sql.Scanner
}

func (w *optionalScanner[T]) Scan(src any) error {
	if src == nil {
		*w.wrapped = optional.Nil[T]()
		return nil
	}

	var value T
	if err := w.createScanner(&value).Scan(src); err != nil {
		*w.wrapped = optional.Optional[T]{}
		return err
	}

	*w.wrapped = optional.Of(value)
	return nil
}

func OptionalValuer[T any](wrapped optional.Optional[T], createValuer func(T) driver.Valuer) *optionalValuer[T] {
	return &optionalValuer[T]{
		wrapped:      wrapped,
		createValuer: createValuer,
	}
}

type optionalValuer[T any] struct {
	wrapped      optional.Optional[T]
	createValuer func(T) driver.Valuer
}

func (w *optionalValuer[T]) Value() (driver.Value, error) {
	if !w.wrapped.IsSet() || w.wrapped.IsNil() {
		return nil, nil
	}

	return w.createValuer(w.wrapped.Value()).Value()
}

func OptionalTexter[T any](wrapped optional.Optional[T], createTexter func(T) modelg.SQLTexter) modelg.SQLTexter {
	return &optionalTexter[T]{
		wrapped:      wrapped,
		createTexter: createTexter,
	}
}

type optionalTexter[T any] struct {
	wrapped      optional.Optional[T]
	createTexter func(T) modelg.SQLTexter
}

func (w *optionalTexter[T]) SQLText(ctx *modelg.SQLTexterContext) (string, error) {
	if !w.wrapped.IsSet() || w.wrapped.IsNil() {
		return "", nil
	}
	return w.createTexter(w.wrapped.Value()).SQLText(ctx)
}
