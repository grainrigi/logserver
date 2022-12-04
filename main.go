package main

import (
	"embed"
	"log"
	"logserver/app"
	"logserver/db"

	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/joho/godotenv"
)

//go:embed migrations/*.sql
var migrations embed.FS

func main() {
	// Load config
	err := godotenv.Load()
	if err != nil {
		// Just warn
		log.Printf("Warning: Could not load .env file: %s", err)
	}

	// DB初期化(embedded FSを添えて)
	d, err := iofs.New(migrations, "migrations")
	if err != nil {
		log.Fatalf("Failed to open migrations file: %s", err)
	}
	if err := db.InitDB("iofs", d); err != nil {
		log.Fatal(err)
	}

	app.Run()
}
