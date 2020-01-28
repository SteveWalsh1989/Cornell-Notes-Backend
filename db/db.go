package db

import (
	"database/sql"
	"fmt"
	 "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func createConn() {
	// Test database
	db, err = sql.Open("mysql", "root:Ilikefood1@tcp(localhost:3306)/sys")
	check(err)
	defer db.Close()
	err = db.Ping()
	check(err)

}

// Checks for non nil errors and prints
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
