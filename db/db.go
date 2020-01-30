package db

import (
	m "FYP_Proto_Backend/model"
	"context"
	"database/sql"
	"fmt"
	"log"
)

var ctx context.Context

//SetupDB ... creates db tables and enters sample data
func SetupDB() {
	db := CreateConn()

	//dropTables(db)    // drop previous tables
	createTables(db)  // create new tables
	addSampleData(db) // insert sample data

	CloseConn(db)

}

//CreateConn ... crates connection with database -returns pointer to db
func CreateConn() *sql.DB {
	// Test database
	db, err := sql.Open("mysql", "root:Ilikefood1@tcp(localhost:3306)/sys")
	check(err)
	err = db.Ping()
	check(err)

	fmt.Println("Connected to db")

	return db
}

//CloseConn ... closes the database connection
func CloseConn(db *sql.DB) {
	db.Close()
	fmt.Println("Disconnected from db")

}

//CreateTables ... create tables in db
func createTables(db *sql.DB) {

	for _, v := range CreateTableCommands {
		insert, err := db.Query(v)
		check(err)
		defer insert.Close()
		// could alsp use 'db.Prepare(v)' along with the follwoing to get more info from rows
		// r, err := insert.Exec()
		// check(err)
		// n, err := r.RowsAffected()
		// check(err)
		fmt.Println("-- Tables Created")
	}
}

//dropTables .. for dev purposes only
func dropTables(db *sql.DB) {
	for k := range CreateTableCommands {
		query := "DROP TABLE " + k
		res, err := db.Query(query)
		check(err)
		defer res.Close()
		fmt.Println("-- Tables Dropped")
	}
}

//AddSampleData ... insert sample data to db
func addSampleData(db *sql.DB) {

	for _, v := range CreateTableCommands {
		insert, err := db.Query(v)
		check(err)
		defer insert.Close()
		// could alsp use 'db.Prepare(v)' along with the follwoing to get more info from rows
		// r, err := insert.Exec()
		// check(err)
		// n, err := r.RowsAffected()
		// check(err)
		fmt.Println("-- Sample Data added to DB")
	}
}

//GetFolder ... using folder ID returns folder from db as Folder struct
func GetFolder(id string) m.Folder {
	db := CreateConn()
	var folder m.Folder

	rows, err := db.Query("SELECT * FROM Folder WHERE id=?", id)
	if err != nil {
		fmt.Println("Error1")
		log.Fatal(err)
	}
	fmt.Println("Rows: ", rows)

	for rows.Next() {
		fmt.Println("here")
		if err := rows.Scan(&folder.Name, &folder.ID, &folder.Status,
			&folder.ItemID, &folder.ItemType, &folder.DateCreated,
			&folder.DateEdited, &folder.DateDeleted); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			fmt.Println("Error")

			log.Fatal(err)
		}
		fmt.Println("here1")

		fmt.Println("folder: ", folder.Name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("here2")

	return folder
}

// Checks for non nil errors and prints
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
