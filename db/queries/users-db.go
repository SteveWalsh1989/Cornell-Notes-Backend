package queries

import (
	"FYP_Proto_Backend/db"
	m "FYP_Proto_Backend/model"
	"fmt"
)

// CreateUser ... adds new user entry to DB
func CreateUser(user m.User) {

	fmt.Println("User from queries: ", user)

	conn := db.CreateConn()

	stmt, err := conn.Prepare("INSERT INTO users (id, name, password, date_created, date_edited) VALUES(?,?,?,?,?);")
	db.Check(err)

	_, errr := stmt.Exec(user.ID, user.Name, user.Password, user.DateCreated, user.DateEdited)

	db.Check(errr)

	db.CloseConn(conn)
}
