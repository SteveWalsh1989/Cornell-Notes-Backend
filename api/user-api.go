package api

import (
	"FYP_Proto_Backend/db"
	q "FYP_Proto_Backend/db/queries"
	m "FYP_Proto_Backend/model"
	"fmt"

	"encoding/json"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

// RegisterUser : registers new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {

	user := m.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	db.Check(err)

	// check if there is an existing user by email and return error if true
	userExistis := q.CheckUserExists(user)
	if userExistis {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("409 - User Email Already Used"))
		json.NewEncoder(w).Encode(http.StatusConflict)
		return
	}
	// set up other user fields
	id, err2 := uuid.NewV4() // create new UUID for new user

	db.Check(err2)
	user.ID = id.String()
	user.DateCreated = time.Now()
	user.DateEdited = time.Now()
	//fmt.Println("time: ", user.DateCreated) // testing - print user details

	q.CreateUser(user) // Call db query

	//fmt.Println(user) // testing - print user details
	json.NewEncoder(w).Encode(&user)
}

// LoginUser : login as exisitng user
func LoginUser(w http.ResponseWriter, r *http.Request) {
	user := m.User{}
	var email []string
	var password []string

	// Get query parameters
	email, _ = r.URL.Query()["email"]
	password, _ = r.URL.Query()["password"]

	user.Email = email[0]
	user.Password = password[0]
	// check if there is an existing user by email and return error if true
	userExistis := q.CheckUserExists(user)
	fmt.Println("1: userExistis: ", userExistis) // testing - print user details
	if userExistis != true {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("404 - Email not found"))
		json.NewEncoder(w).Encode(http.StatusBadRequest)
		return
	}

	loggedInUser := q.LoginUser(user)           // Get user details
	fmt.Println("2: email: ", user.Email)       // testing - print user details
	fmt.Println("2: password: ", user.Password) // testing - print user details

	//fmt.Println(user) // testing - print user details
	json.NewEncoder(w).Encode(&loggedInUser)
}
