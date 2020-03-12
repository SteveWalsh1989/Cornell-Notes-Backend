package queries

import (
	"FYP_Backend/db"
	m "FYP_Backend/model"
	"fmt"
)

// CreateUser ... adds new user entry to DB
func CreateUser(user m.User) {

	//fmt.Println("User from queries: ", user)

	conn := db.CreateConn()

	stmt, err := conn.Prepare("INSERT INTO users (id, user_name, email, password, date_created, date_edited) VALUES(?,?,?,?,?,?);")
	db.Check(err)

	_, errr := stmt.Exec(user.ID, user.Name, user.Email, user.Password, user.DateCreated, user.DateEdited)

	db.Check(errr)

	db.CloseConn(conn)
}

//CheckUserExists ... uses the email provided to check if the user already has an account
func CheckUserExists(user m.User) bool {
	res := false
	conn := db.CreateConn()
	var existingUser m.User // stores if there is an existing user

	query := "SELECT email FROM Users WHERE email= '" + user.Email + "'"
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
	return res
}

//LoginUser ... logs in existing user
func LoginUser(user m.User) m.User {
	conn := db.CreateConn()
	//worked
	fmt.Println("LoginUser: email: ", user.Email)       // testing - print user details
	fmt.Println("LoginUser: password: ", user.Password) // testing - print user details
	var loggedInUser m.User                             // stores if there is an existing user

	query := "SELECT u.id, u.email, u.user_name FROM Users u WHERE u.email= '" + user.Email + "' AND u.password= '" + user.Password + "'"
	rows, err := conn.Query(query)
	db.Check(err)
	for rows.Next() {
		if err := rows.Scan(&loggedInUser.ID, &loggedInUser.Email, &loggedInUser.Name); err != nil {
			fmt.Println("Error Scanning Rows: ", err)
		}
		//worked
		fmt.Println("loggedInUser: email: ", loggedInUser.Email)       // testing - print user details
		fmt.Println("loggedInUser: password: ", loggedInUser.Password) // testing - print user details

	}
	err = rows.Err()
	db.Check(err)
	db.CloseConn(conn)

	return loggedInUser
}
