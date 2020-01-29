package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

var ctx context.Context
var tables = map[string]string{
	"Customer": "CREATE TABLE customer (name VARCHAR(20));",
}

func CreateConn() *sql.DB {
	// Test database
	db, err := sql.Open("mysql", "root:Ilikefood1@tcp(localhost:3306)/sys")
	check(err)
	err = db.Ping()
	check(err)

	return db
}

func CloseConn(db *sql.DB) {
	db.Close()
}

func CreateTables(db *sql.DB) {
	fmt.Println("table: ", tables)

	for k, v := range tables {
		fmt.Println("Key: ", k)
		fmt.Println("Value: ", v)

		stmt, err := db.Prepare(v)
		check(err)
		defer stmt.Close()

		r, err := stmt.Exec()
		check(err)

		n, err := r.RowsAffected()
		check(err)

		fmt.Println("CREATED TABLE:", k, ",Rows", n)

	}

}

func GetFolder(db *sql.DB) {
	fmt.Println("Reached 1")
	var (
		id   int
		name string
	)
	rows, err := db.Query("select * from Folder;")
	fmt.Println("Reached 2")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Reached 3")

	defer rows.Close()
	fmt.Println("Reached 4")

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
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
