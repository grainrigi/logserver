package db

import (
	"context"
	"logserver/data"
)

func InsertLog(ctx context.Context, l *data.Log) error {
	_, err := db.NewInsert().Model(l).Exec(ctx)
	return err
}

func GetLogs(ctx context.Context) ([]data.Log, error) {
	var logs []data.Log
	err := db.NewSelect().Model(&logs).Relation("Operator").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return logs, nil
}
