package db

import (
	m "FYP_Proto_Backend/model"
	"context"
	"database/sql"
	"fmt"
)

var ctx context.Context

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

//CreateConn ... crates connection with database -returns pointer to db
func CreateConn() *sql.DB {
	// Test database
	db, err := sql.Open("mysql", "root:Ilikefood1@tcp(localhost:3306)/sys")
	Check(err)
	err = db.Ping()
	Check(err)
	LogDBConn("DB Connected")
	return db
}

//CloseConn ... closes the database connection
func CloseConn(db *sql.DB) {
	db.Close()
	LogDBConn("DB Disconnected")

}

//CreateTables ... create tables in db
func createTables(db *sql.DB) {
	LogTitle("Creating DB")
	numCreated := 0
	for k, v := range CreateTableCommands {
		insert, err := db.Query(v)
		Check(err)
		defer insert.Close()
		// could alsp use 'db.Prepare(v)' along with the follwoing to get more info from rows
		// r, err := insert.Exec()
		// check(err)
		// n, err := r.RowsAffected()
		// check(err)
		fmt.Println("-- Table Created: ", k)
		numCreated++
	}

	fmt.Println("-- Tables Created:", numCreated)

}

//dropTables .. for dev purposes only
func dropTables(db *sql.DB) {
	for k := range CreateTableCommands {
		query := "DROP TABLE " + k
		res, err := db.Query(query)
		Check(err)
		defer res.Close()
	}
	LogTitle("Old Tables Dropped")

}

//AddSampleData ... insert sample data to db
func addSampleData(db *sql.DB) {
	for _, v := range InsertSampleDataCommands {
		insert, err := db.Query(v)
		Check(err)
		defer insert.Close()
	}
	LogTitle("Sample Data added to DB")
}

//GetFolder ... using folder ID returns folder from db as Folder struct
func GetFolder(id string) m.Folder {
	db := CreateConn()
	var folder m.Folder

	rows, err := db.Query("SELECT * FROM Folders WHERE id=?", id)
	Check(err)
	for rows.Next() {
		if err := rows.Scan(&folder.Name, &folder.ID, &folder.Status,
			&folder.DateCreated, &folder.DateEdited); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			fmt.Println("Error", err)
		}
		//fmt.Println("folder: ", folder.Name)
	}
	err = rows.Err()
	Check(err)

	return folder
}

//GetFolders ... using folder ID returns folder from db as Folder struct
func GetFolders() []m.Folder {
	db := CreateConn()
	var folder m.Folder
	var folders = []m.Folder{}

	rows, err := db.Query("SELECT * FROM Folders")
	Check(err)

	for rows.Next() {
		if err := rows.Scan(&folder.ID, &folder.Name, &folder.Status,
			&folder.DateCreated, &folder.DateEdited); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			fmt.Println("Error: ", err)
		}
		// fmt.Println("folder: ", folder.Name)
		folders = append(folders, folder)
	}
	err = rows.Err()
	Check(err)
	return folders
}
