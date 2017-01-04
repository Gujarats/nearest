package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var Postgre *sqlx.DB

func Connect() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=potgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	Postgre = db
}
