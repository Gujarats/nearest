package database

//
//import (
//	"database/sql"
//	_ "github.com/lib/pq"
//	"log"
//)
//
//var PostgreDb *sql.DB
//
//func InitDb() {
//	db, err := sql.Open("postgres", "user=techacademy password=123qwe!@#QWE dbname=tokopedia-talk host=192.168.100.126 port=5432 sslmode=disable")
//
//	if err != nil {
//		log.Fatal("Error : data source arument not valid")
//	}
//
//	err = db.Ping()
//	if err != nil {
//		log.Fatal("Error: Could not establish a connection with the database")
//	}
//
//	PostgreDb = db
//
//}
