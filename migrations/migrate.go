package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rhmdnd/compserv-poc/compserv"
)

func main() {
	c := compserv.ParseConfig()
	connStr := compserv.GetDatabaseConnectionString(c)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Cannot connect to database: %s", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		c["db_name"], driver)
	if err != nil {
		log.Fatalf("Unable to load migrations: %s", err)
	}
	m.Up()
}
