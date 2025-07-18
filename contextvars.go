package modelg

import (
	"context"
)

type contextVarsKeyType struct{}

var contextVarsKey = contextVarsKeyType{}

type ContextVars interface {
	GetScanTarget(name string) any
}

type nullContextVars struct{}

func (n nullContextVars) GetScanTarget(name string) any {
	return nil
}

func ContextVarsFromContext(ctx context.Context) ContextVars {
	vars, ok := ctx.Value(contextVarsKey).(ContextVars)
	if !ok {
		return nullContextVars{}
	}
	return vars
}

func WithContextVars(ctx context.Context, vars ContextVars) context.Context {
	return context.WithValue(ctx, contextVarsKey, vars)
}
