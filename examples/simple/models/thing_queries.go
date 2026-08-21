package models

import "context"

func (q *ThingQueries) EnsureTable(ctx context.Context) error {
	if err := q.dropTable(ctx); err != nil {
		return err
	}
	if err := q.createTable(ctx); err != nil {
		return err
	}
	return nil
}
