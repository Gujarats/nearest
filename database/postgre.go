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
	Password string
}

type PostgreSystem interface {
	Init()
	Connect() (*sqlx.DB, error)
}

var logger *log.Logger

func (p PostgreHost) Init() {
	logger = log.New(os.Stderr,
		"Postgre",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func (p *PostgreHost) Connect() (*sqlx.DB, error) {
	connection := fmt.Sprintf("user=%v password= %v dbname=%v sslmode=%v", p.Username, p.Password, p.Database, p.Ssl)
	db, err := sqlx.Connect(
		p.Driver,
		connection)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	return db, nil
}

func GetPostgreDb(postgre PostgreSystem) (*sqlx.DB, error) {
	return postgre.Connect()
}
