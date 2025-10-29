package main

import (
	"context"
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5"
)

const dbURL = "postgres://login:pass@localhost:5432/db-name"

var ctx = context.Background()

func main() {
	goquInsert()
	// goquSelect()
	// goquUpdate()
	// goquDelete()
	// Seeder()
}

func goquInsert() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	ds := goqu.From("users").
		Insert().
		Rows(
			goqu.Record{"name": "Alice", "age": 20},
			goqu.Record{"name": "Bob", "age": 30},
			goqu.Record{"name": "Charlie", "age": 40},
		)

	sql, _, err := ds.ToSQL()
	if err != nil {
		log.Fatalf("ds.ToSQL: %v\n", err)
	}

	fmt.Println(sql)

	tag, err := conn.Exec(ctx, sql)
	if err != nil {
		log.Fatalf("conn.Exec: %v\n", err)
	}

	fmt.Printf("Tag: %v\n", tag)
}

func goquSelect() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	ds := goqu.From("users").
		Select("name", "age").
		Where(
			goqu.And(
				goqu.C("age").Gte(20),
				goqu.C("name").Eq("Alice"),
			),
		).
		Order(
			goqu.I("name").Desc(),
		)

	sql, _, err := ds.ToSQL()
	if err != nil {
		log.Fatalf("ds.ToSQL: %v\n", err)
	}

	fmt.Println(sql)

	rows, err := conn.Query(ctx, sql)
	if err != nil {
		log.Fatalf("conn.Query: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var age int

		err := rows.Scan(&name, &age)
		if err != nil {
			log.Fatalf("rows.Scan: %v\n", err)
		}

		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}
}

func goquUpdate() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	ds := goqu.From("users").
		Update().
		Set(
			goqu.Record{"age": 33},
		).
		Where(
			goqu.C("name").Eq("Alice"),
		)

	sql, _, err := ds.ToSQL()
	if err != nil {
		log.Fatalf("ds.ToSQL: %v\n", err)
	}

	fmt.Println(sql)

	tag, err := conn.Exec(ctx, sql)
	if err != nil {
		log.Fatalf("conn.Exec: %v\n", err)
	}

	fmt.Printf("Tag: %v\n", tag)
}

func goquDelete() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	ds := goqu.From("users").
		Delete().
		Where(
			goqu.C("name").Eq("Alice"),
		)

	sql, _, err := ds.ToSQL()
	if err != nil {
		log.Fatalf("ds.ToSQL: %v\n", err)
	}

	fmt.Println(sql)

	tag, err := conn.Exec(ctx, sql)
	if err != nil {
		log.Fatalf("conn.Exec: %v\n", err)
	}

	fmt.Printf("Tag: %v\n", tag)
}

func Seeder() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	records := make([]any, 0, 1000)

	for range 1000 {
		records = append(records, goqu.Record{
			"name": gofakeit.Name(),
			"age":  gofakeit.IntRange(18, 120),
		})
	}

	ds := goqu.Insert("users").Rows(records...)

	sql, _, err := ds.ToSQL()
	if err != nil {
		log.Fatalf("ds.ToSQL: %v\n", err)
	}

	fmt.Println(sql)

	tag, err := conn.Exec(ctx, sql)
	if err != nil {
		log.Fatalf("conn.Exec: %v\n", err)
	}

	fmt.Printf("Tag: %v\n", tag)
}
