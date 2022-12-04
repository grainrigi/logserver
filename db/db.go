package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	db *bun.DB
)

// 初期化処理
func InitDB(sourceName string, sourceInstance source.Driver) error {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		return fmt.Errorf("DB_URL is empty. Please specify database connection URL")
	}

	db = newDB(dsn)

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.FromEnv("BUNDEBUG"),
	))

	// マイグレーションを実行
	m, err := migrate.NewWithSourceInstance(
		sourceName,
		sourceInstance,
		dsn,
	)
	if err != nil {
		return fmt.Errorf("Failed to migrate DB: %s", err)
	}
	m.Up()

	return nil
}

// DBへの接続
func newDB(dsn string) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	return db
}
