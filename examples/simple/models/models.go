package models

import (
	"context"

	"github.com/partite-ai/modelg"
	"github.com/partite-ai/optional"
)

//go:generate go run github.com/partite-ai/modelg/cmd/modelg

type Thing struct {
	ID     int64 `modelg:"pk,computed"`
	Name   string
	Status int
}

type thingQueries interface {
	dropTable(ctx context.Context) error
	createTable(ctx context.Context) error
	FindBy(ctx context.Context, filter *modelg.DynamicFilters, order optional.Optional[OrderBy]) ([]*Thing, error)
	FindCount(ctx context.Context, filter *modelg.DynamicFilters) (int, error)
	FindByNameAndStatus(ctx context.Context, name optional.Optional[string], status optional.Optional[int]) ([]*Thing, error)
	Delete(ctx context.Context, thingID int64) error
	FindActiveByNames(ctx context.Context, names modelg.InValues) ([]*Thing, error)
}
