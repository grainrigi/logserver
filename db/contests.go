package db

import (
	"context"
	"logserver/data"
)

func ReadContests(ctx context.Context) ([]data.Contest, error) {
	cs := make([]data.Contest, 0)

	err := db.NewSelect().Model(&cs).Scan(ctx)

	return cs, err
}

func ReadContest(ctx context.Context, id int) (data.Contest, error) {
	c := data.Contest{ID: id}

	err := db.NewSelect().Model(&c).WherePK().Scan(ctx)

	return c, err
}

func InsertContest(ctx context.Context, c *data.Contest) (int, error) {
	var rc data.Contest
	c.ID = 0

	_, err := db.NewInsert().Model(c).Returning("id").Exec(ctx, &rc)

	return rc.ID, err
}

func DeleteContest(ctx context.Context, id int) error {
	c := data.Contest{ID: id}

	_, err := db.NewDelete().Model(&c).WherePK().Exec(ctx)

	return err
}

func UpdateContest(ctx context.Context, c *data.Contest) error {
	_, err := db.NewUpdate().Model(c).WherePK().Exec(ctx)

	return err
}
