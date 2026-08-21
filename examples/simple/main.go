package main

import (
	"context"
	"fmt"
	"os"

	"github.com/partite-ai/modelg"
	"github.com/partite-ai/modelg/examples/simple/models"
	"github.com/partite-ai/optional"
	"zombiezen.com/go/sqlite/sqlitex"
)

func main() {
	ctx := context.Background()
	pool, err := sqlitex.NewPool("file:simple.db", sqlitex.PoolOptions{
		PoolSize: 1,
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create SQLite connection pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	db := modelg.NewSqliteDB(pool)

	queries := models.NewThingQueries(db)
	if err := queries.EnsureTable(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create table: %v\n", err)
	}

	check(queries.CreateThing(ctx, &models.ThingCreateParams{
		Name:   "thing1",
		Status: 1,
	}))
	check(queries.CreateThing(ctx, &models.ThingCreateParams{
		Name:   "thing2",
		Status: 1,
	}))
	check(queries.CreateThing(ctx, &models.ThingCreateParams{
		Name:   "thing3",
		Status: 2,
	}))

	filter := modelg.DynamicFiltersAnd()
	filter.AddEq("name", "thing1")
	filter.AddEq("status", 1)
	for _, thing := range check(queries.FindBy(ctx, filter, optional.Nil[models.OrderBy]())) {
		fmt.Fprintf(os.Stdout, "FindBy found: %s\n", thing.Name)
	}

	for _, thing := range check(queries.FindBy(ctx, nil, optional.Of(models.OrderByDesc))) {
		fmt.Fprintf(os.Stdout, "FindBy desc found: %s\n", thing.Name)
	}

	for _, thing := range check(queries.FindByNameAndStatus(ctx, optional.Of("thing2"), optional.Optional[int]{})) {
		fmt.Fprintf(os.Stdout, "FindByNameAndStatus found: %s\n", thing.Name)
	}

	for _, thing := range check(queries.FindActiveByNames(ctx, modelg.InValues{"thing1", "thing3"})) {
		fmt.Fprintf(os.Stdout, "FindActiveByNames 2 values found: %s\n", thing.Name)
	}

	for _, thing := range check(queries.FindActiveByNames(ctx, modelg.InValues{})) {
		fmt.Fprintf(os.Stdout, "FindActiveByNames no values found: %s\n", thing.Name)
	}
}

func check[T any](v T, err error) T {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	return v
}
