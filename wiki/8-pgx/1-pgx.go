package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const dbURL = "postgres://login:pass@localhost:5432/db-name"

var ctx = context.Background()

func main() {
	// DropTable()
	CreateTable()
	// CreateIndex()
	// InsertOne()
	// InsertMany()
	// Update()
	// Delete()
	// SelectAll()
	// SelectOne()
	// Transaction()
	// TransactionSerializable()
}

func CreateTable() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("conn.Ping: %v\n", err)
	}

	const sql = `CREATE TABLE IF NOT EXISTS users (
					id SERIAL PRIMARY KEY,
					name TEXT NOT NULL,
					age INT
				);`

	tag, err := conn.Exec(ctx, sql)
	if err != nil {
		log.Fatalf("conn.Exec: %v\n", err)
	}

	fmt.Printf("Tag: %v\n", tag)
}

func DropTable() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	const sql = `DROP TABLE IF EXISTS users;`

	tag, err := conn.Exec(ctx, sql)
	if err != nil {
		log.Fatalf("conn.Exec: %v\n", err)
	}

	fmt.Printf("Tag: %v\n", tag)
}

func CreateIndex() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	const sql = `CREATE INDEX IF NOT EXISTS idx_users_name ON users (name);`

	tag, err := conn.Exec(ctx, sql)
	if err != nil {
		log.Fatalf("conn.Exec: %v\n", err)
	}

	fmt.Printf("Tag: %v\n", tag)
}

func InsertOne() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	const sql = `INSERT INTO users (name, age) VALUES ($1, $2);`

	tag, err := conn.Exec(ctx, sql, "Alice", 20)
	if err != nil {
		log.Fatalf("conn.Exec: %v\n", err)
	}

	fmt.Printf("Tag: %v\n", tag)
}

func InsertMany() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	sql := `INSERT INTO users (name, age) VALUES `

	var args []any

	args = append(args, "Alice", 42)
	sql += "($1, $2),"

	args = append(args, "Bob", 25)
	sql += "($3, $4),"

	args = append(args, "Charlie", 30)
	sql += "($5, $6);"

	fmt.Println(sql)

	tag, err := conn.Exec(ctx, sql, args...)
	if err != nil {
		log.Fatalf("conn.Exec: %v\n", err)
	}

	fmt.Printf("Tag: %v\n", tag)
}

func Update() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	const sql = `UPDATE users SET name = $1, age = $2 WHERE id = $3;`

	tag, err := conn.Exec(ctx, sql, "Angela", 18, 1)
	if err != nil {
		log.Fatalf("conn.Exec: %v\n", err)
	}

	fmt.Printf("Tag: %v\n", tag)
}

func Delete() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	const sql = `DELETE FROM users WHERE id = $1;`

	tag, err := conn.Exec(ctx, sql, 3)
	if err != nil {
		log.Fatalf("conn.Exec: %v\n", err)
	}

	fmt.Printf("Tag: %v\n", tag)
}

func SelectAll() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	const sql = `SELECT * FROM users;`

	rows, err := conn.Query(ctx, sql)
	if err != nil {
		log.Fatalf("conn.Query: %v\n", err)
	}

	for rows.Next() {
		var id int
		var name string
		var age int

		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatalf("rows.Scan: %v\n", err)
		}

		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}

func SelectOne() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	const sql = `SELECT * FROM users WHERE id = $1;`

	row := conn.QueryRow(ctx, sql, 1)

	var id pgtype.Int8
	var name pgtype.Text
	var age pgtype.Int8

	err = row.Scan(&id, &name, &age)
	if err != nil {
		log.Fatalf("row.Scan: %v\n", err)
	}

	fmt.Printf("ID: %d, Name: %s, Age: %d\n", id.Int64, name.String, age.Int64)
}

func Transaction() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Fatalf("conn.Begin: %v\n", err)
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			log.Fatalf("tx.Rollback: %v\n", err)
		}
	}()

	_, err = tx.Exec(ctx, `INSERT INTO users (name, age) VALUES ($1, $2);`, "Transaction_1", 42)
	if err != nil {
		log.Fatalf("tx.Exec: %v\n", err)
	}

	_, err = tx.Exec(ctx, `INSERT INTO users (name, age) VALUES ($1, $2);`, "Transaction_2", 24)
	if err != nil {
		log.Fatalf("tx.Exec: %v\n", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalf("tx.Commit: %v\n", err)
	}
}

func TransactionSerializable() {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("pgx.Connect: %v\n", err)
	}

	txOptions := pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	}

	tx, err := conn.BeginTx(ctx, txOptions)
	if err != nil {
		log.Fatalf("conn.BeginTx: %v\n", err)
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			log.Fatalf("tx.Rollback: %v\n", err)
		}
	}()

	_, err = tx.Exec(ctx, `INSERT INTO users (name, age) VALUES ($1, $2);`, "Serializable_1", 42)
	if err != nil {
		log.Fatalf("tx.Exec: %v\n", err)
	}

	_, err = tx.Exec(ctx, `INSERT INTO users (name, age) VALUES ($1, $2);`, "Serializable_2", 24)
	if err != nil {
		log.Fatalf("tx.Exec: %v\n", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalf("tx.Commit: %v\n", err)
	}
}
