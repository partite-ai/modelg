package modelg

import (
	"context"
)

type SimpleSQLTexter interface {
	SQLText() string
}

func SQLTexterForSimpleSQLTexter(txter SimpleSQLTexter) SQLTexter {
	return simpleSqlTexterWrapper{txter}
}

type simpleSqlTexterWrapper struct {
	s SimpleSQLTexter
}

func (w simpleSqlTexterWrapper) SQLText(ctx *SQLTexterContext) (string, error) {
	return w.s.SQLText(), nil
}

type SimpleModeSQLTexter interface {
	SQLText(mode string) string
}

func SQLTexterForSimpleModeSQLTexter(txter SimpleModeSQLTexter) SQLTexter {
	return simpleModeSqlTexterWrapper{txter}
}

type simpleModeSqlTexterWrapper struct {
	s SimpleModeSQLTexter
}

func (w simpleModeSqlTexterWrapper) SQLText(ctx *SQLTexterContext) (string, error) {
	return w.s.SQLText(ctx.Mode), nil
}

type SQLTexter interface {
	SQLText(ctx *SQLTexterContext) (string, error)
}

type SQLTexterContext struct {
	Context context.Context
	Vars    QueryVariablesScope
	Mode    string
}
