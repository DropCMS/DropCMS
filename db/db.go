package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


var Db *sql.DB
var err error
func ConnectDb() {
	dsn := "root:Arpit@tcp(localhost:3306)/arpit"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
}
