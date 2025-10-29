package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	const (
		filePath = "file://migration"
		dbURL    = "postgres://login:pass@localhost:5432/db-name?sslmode=disable"
	)

	m, err := migrate.New(filePath, dbURL)
	if err != nil {
		log.Fatal(err)
	}

	{ // Применить все миграции
		err = m.Up()
		if err != nil {
			log.Fatal(err)
		}
	}

	{ // Откатить все миграции
		err = m.Down()
		if err != nil {
			log.Fatal(err)
		}
	}

	//{ // Узнать версию текущей миграции
	//	version, dirty, err := m.Version()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	log.Printf("Version: %v, Dirty: %v\n", version, dirty)
	//}

	//{ // Применить миграцию до определенной версии
	//	err = m.Migrate(20250213144246)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	{ // Применить миграцию до определенной версии при dirty состоянии
		err = m.Force(20250213144246)
		if err != nil {
			log.Fatal(err)
		}
	}
}
