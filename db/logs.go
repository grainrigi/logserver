package db

import (
	"context"
	"logserver/data"
)

func ReadLogs(ctx context.Context, cid int) ([]data.Log, error) {
	var logs []data.Log
	err := db.NewSelect().Model(&logs).Relation("Operator").Where("contest = ?", cid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func InsertLog(ctx context.Context, l *data.Log) (int, error) {
	rl := data.Log{}
	_, err := db.NewInsert().Model(l).Returning("id").Exec(ctx, &rl)
	return rl.ID, err
}

func DeleteLog(ctx context.Context, id int) error {
	l := data.Log{ID: id}
	_, err := db.NewDelete().Model(&l).WherePK().Exec(ctx)
	return err
}

func UpdateLog(ctx context.Context, l *data.Log) error {
	_, err := db.NewUpdate().Model(l).WherePK().Exec(ctx)

	return err
}
