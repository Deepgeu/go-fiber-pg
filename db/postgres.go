package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool
var Ctx = context.Background()

func ConnectDB() {
	dsn := "postgres://postgres:1445632@localhost:5432/redis" // postgres password

	var err error
	DB, err = pgxpool.New(Ctx, dsn)
	if err != nil {
		log.Fatal("//// PostgreSQL connection error:", err)
	}

	_, err = DB.Exec(Ctx, `
		CREATE TABLE IF NOT EXISTS records (
			id SERIAL PRIMARY KEY,
			key TEXT UNIQUE,
			value TEXT
		)
	`)
	if err != nil {
		log.Fatal("?? Table creation failed:", err)
	}

	fmt.Println("// PostgreSQL connected and table ready!")
}
