package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

var ctx context.Context
var dbCommands = map[string]string{
	// "createUser": "CREATE TABLE user (name VARCHAR(20));",
	"insertUser": "INSERT INTO user (name) VALUES('mikes');",
}

func CreateConn() *sql.DB {
	// Test database
	db, err := sql.Open("mysql", "root:Ilikefood1@tcp(localhost:3306)/sys")
	check(err)
	err = db.Ping()
	check(err)

	fmt.Println("Connected to db")

	return db
}

func CloseConn(db *sql.DB) {
	db.Close()
	fmt.Println("Disconnected from db")

}

func CreateTables(db *sql.DB) {
	fmt.Println("dbCommands: ", dbCommands)

	for _, v := range dbCommands {
		insert, err := db.Query(v)
		check(err)
		defer insert.Close()
		// could alsp use 'db.Prepare(v)' along with the follwoing to get more info from rows
		// r, err := insert.Exec()
		// check(err)
		// n, err := r.RowsAffected()
		// check(err)
		fmt.Println("Command send successfully")
	}

}

func GetFolder(db *sql.DB) {
	var name string
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		// DO shit here to assign values
		fmt.Println("Name: ", name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

// Checks for non nil errors and prints
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
