package db

import (
	"context"
	"database/sql"
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
