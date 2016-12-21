package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var MysqlDb *sql.DB

func InitMysqlDb() {
	db, err := sql.Open("mysql", "root:root@/tokopediaDb")

	if err != nil {
		log.Fatal("Error : Opening database data source argument not valid")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error could not established connection with database")
	}

	//database succesfully connected
	MysqlDb = db
}
