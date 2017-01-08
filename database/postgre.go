package database

// this file is used for connecting to postgre database system

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreHost struct {
	Driver   string
	Database string
	Username string
	Ssl      string
	Passowrd string
}

type PostgreSystem interface {
	Init()
	Connect() (*sqlx.DB, error)
}

var logger *log.Logger

func (self PostgreHost) Init() {
	logger = log.New(os.Stderr,
		"Postgre",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func (self *PostgreHost) Connect() (*sqlx.DB, error) {
	connection := fmt.Sprintf("user=%v password= %v dbname=%v sslmode=%v", self.Username, self.Passowrd, self.Database, self.Ssl)
	db, err := sqlx.Connect(
		self.Driver,
		connection)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	return db, nil
}
