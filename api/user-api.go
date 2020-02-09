package api

import (
	"FYP_Proto_Backend/db"
	q "FYP_Proto_Backend/db/queries"
	m "FYP_Proto_Backend/model"

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
		json.NewEncoder(w).Encode()
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
