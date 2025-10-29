package main

import (
	"context"
	"fmt"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5"
)

const dbURL = "postgres://login:pass@localhost:5432/db-name"

var ctx = context.Background()

func main() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	// Сама инъекция
	badCode := "'; DROP TABLE profile; --"

	// Инъекция сработает, так никогда не делайте
	sql := fmt.Sprintf("SELECT * FROM profile WHERE name = '%s'", badCode)
	// Output: SELECT * FROM profile WHERE name = ''; DROP TABLE profile; --'

	// Инъекция не сработает, в goqu есть защита
	sql, _, _ = goqu.Select().From("profile").
		Where(goqu.Ex{"name": badCode}).ToSQL()
	// Output: SELECT * FROM "profile" WHERE ("name" = '''; DROP TABLE profile; --')

	// Инъекция не сработает, если вы передаёте параметры через плейсхолдеры $1, $2...
	sql = "SELECT * FROM profile WHERE name = $1"

	_, err = conn.Exec(ctx, sql, badCode)
	if err != nil {
		log.Fatalf("Error executing SQL: %v", err)
	}
}
