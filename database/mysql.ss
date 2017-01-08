//package database
//
//// this file is used for connecting to mysql database system
//
//import (
//	"database/sql"
//	"fmt"
//	"log"
//	"os"
//
//	_ "github.com/go-sql-driver/mysql"
//)
//
//type MysqlHost struct {
//	Driver   string
//	Username string
//	Password string
//	Database string
//}
//
//type MysqlSystem interface {
//	Init()
//	Connect() (*sql.DB, error)
//}
//
//// we use logger here to prefix the log and display the error line
//var logger *log.Logger
//
//// using Mysql word as the prefix of the logger
//func (self MysqlHost) Init() {
//	logger = log.New(os.Stderr,
//		"MySql",
//		log.Ldate|log.Ltime|log.Lshortfile)
//}
//
//func (self MysqlHost) Connect() (*sql.DB, error) {
//	connection := fmt.Sprintf("%v:%v@/%v", self.Username, self.Passowrd, self.Database)
//	db, err := sql.Open(self.Driver, connection)
//
//	if err != nil {
//		logger.Fatal(err)
//		return nil, err
//	}
//
//	err = db.Ping()
//	if err != nil {
//		logger.Fatal(err)
//		return nil, err
//	}
//
//	//database succesfully connected
//	return db, nil
//}
