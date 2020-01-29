package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

var ctx context.Context
var dbCommands = map[string]string{
	"createUser": "CREATE TABLE user (name VARCHAR(20));",
	"insertUser": "INSERT INTO user (name) VALUES('steve');",
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
}

func CreateTables(db *sql.DB) {
	fmt.Println("dbCommands: ", dbCommands)

	for k, v := range dbCommands {
		fmt.Println("Key: ", k)
		fmt.Println("Value: ", v)

		stmt, err := db.Prepare(v)
		check(err)
		defer stmt.Close()

		r, err := stmt.Exec()
		check(err)

		n, err := r.RowsAffected()
		check(err)

		fmt.Println("Command:", k, ",Rows", n)

	}

}

func GetFolder(db *sql.DB) {
	var name string
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		fmt.Println("Reached 5")

		err := rows.Scan(&name)
		if err != nil {
			fmt.Println("Error")
			log.Fatal(err)
		}

		fmt.Println("Name: ", name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Reached End")

}

// Checks for non nil errors and prints
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
