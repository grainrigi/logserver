package db

import (
	"context"
	"logserver/data"
)

func ReadOperators(ctx context.Context) ([]data.Operator, error) {
	ops := make([]data.Operator, 0)

	if err := db.NewSelect().Model(&ops).Scan(ctx); err != nil {
		return nil, err
	}

	return ops, nil
}

func InsertOperator(ctx context.Context, op *data.Operator) error {
	op.ID = 0

	_, err := db.NewInsert().Model(op).Exec(ctx)

	return err
}

func DeleteOperator(ctx context.Context, id int) error {
	op := data.Operator{ID: id}

	_, err := db.NewDelete().Model(&op).WherePK().Exec(ctx)

	return err
}

func UpdateOperator(ctx context.Context, op *data.Operator) error {
	_, err := db.NewUpdate().Model(op).WherePK().Exec(ctx)

	return err
}
