package queries

import (
	"FYP_Proto_Backend/db"
	m "FYP_Proto_Backend/model"
	"fmt"
)

// CreateUser ... adds new user entry to DB
func CreateUser(user m.User) {

	//fmt.Println("User from queries: ", user)

	conn := db.CreateConn()

	stmt, err := conn.Prepare("INSERT INTO users (id, name, email, password, date_created, date_edited) VALUES(?,?,?,?,?,?);")
	db.Check(err)

	_, errr := stmt.Exec(user.ID, user.Name, user.Email, user.Password, user.DateCreated, user.DateEdited)

	db.Check(errr)

	db.CloseConn(conn)
}

//CheckUserExists ... uses the email provided to check if the user already has an account
func CheckUserExists(user m.User) bool {
	res := false

	db.LogValue("user email: ", user.Email)

	conn := db.CreateConn()
	var existingUser m.User // stores if there is an existing user

	query := "SELECT email FROM Users WHERE email= '" + user.Email + "'"

	fmt.Println("Query: ", query)

	rows, err := conn.Query(query)
	db.Check(err)

	for rows.Next() {
		if err := rows.Scan(&existingUser.Email); err != nil {
			fmt.Println("Error Scanning Rows: ", err)
		}
		//fmt.Println("Existing email Found: ", existingUser.Email)
		res = true
	}
	err = rows.Err()
	db.Check(err)

	db.CloseConn(conn)

	fmt.Println("res:", res)

	return res
}
