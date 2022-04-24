package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect()(*sql.DB) {

	dsn := "root:@tcp(localhost:3306)/testGO"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}
