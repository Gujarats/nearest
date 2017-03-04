package database

// this file is used for connecting to postgre database system

import (
	"fmt"

	"github.com/Gujarats/API-Golang/util/logger"
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

func init() {
	logger.InitLogger("POSTGRE  :: ", "../logs", "Postgre.txt")
}

type PostgreSystem interface {
	Init()
	Connect() *sqlx.DB
}

func (p PostgreHost) Init() {

}

func (p *PostgreHost) Connect() *sqlx.DB {
	connection := fmt.Sprintf("user=%v password= %v dbname=%v sslmode=%v", p.Username, p.Password, p.Database, p.Ssl)
	db, err := sqlx.Connect(
		p.Driver,
		connection,
	)

	logger.CheckError("Postgre", err)

	return db
}

func GetPostgreDb(postgre PostgreSystem) *sqlx.DB {
	return postgre.Connect()
}
