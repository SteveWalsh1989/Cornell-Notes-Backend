package db

import (
	"context"
	"database/sql"
	"fmt"
)

var ctx context.Context

//CreateConn ... crates connection with database -returns pointer to db
func CreateConn() *sql.DB {
	// Test database
	db, err := sql.Open("mysql", "root:Ilikefood1@tcp(localhost:3306)/sys?parseTime=true")
	Check(err)
	err = db.Ping()
	Check(err)
	//LogDBConn("DB Connected")
	return db
}

//CloseConn ... closes the database connection
func CloseConn(db *sql.DB) {
	db.Close()
	//LogDBConn("DB Disconnected")

}

//SetupDB ... creates db tables and enters sample data
func SetupDB() {
	LogTitle("Setting up Database")

	db := CreateConn()

	dropTables(db)    // drop previous tables
	createTables(db)  // create new tables
	addSampleData(db) // insert sample data

	CloseConn(db)
	LogTitle("Setup Complete")

}

/*
The following three functions are only ran once when starting program:
	- Drop all existing tables in DB
	- Initalise the DB with required tables for app
	- Insert sample data into DB for testing
*/

//CreateTables ... create tables in db
func createTables(db *sql.DB) {
	LogTitle("Creating DB")
	numCreated := 1
	for k, v := range CreateTables {
		insert, err := db.Query(v)
		Check(err)
		defer insert.Close()
		// could alsp use 'db.Prepare(v)' along with the follwoing to get more info from rows
		// r, err := insert.Exec()
		// check(err)
		// n, err := r.RowsAffected()
		// check(err)
		fmt.Printf("-- %v Table Created: %v \n", numCreated, k)
		numCreated++
	}

}

//dropTables .. for dev purposes only
func dropTables(db *sql.DB) {
	for k := range CreateTables {
		query := "DROP TABLE " + k
		res, err := db.Query(query)
		Check(err)
		defer res.Close()
	}
	LogTitle("Old Tables Dropped")

}

//AddSampleData ... insert sample data to db
func addSampleData(db *sql.DB) {
	for _, v := range InsertSampleData {
		insert, err := db.Query(v)
		Check(err)
		defer insert.Close()
	}
	LogTitle("Sample Data added to DB")
}
