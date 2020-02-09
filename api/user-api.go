package api

import (
	"FYP_Proto_Backend/db"
	m "FYP_Proto_Backend/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

// RegisterUser : registers new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {

	// Check form data correct

	// create new UUID for new user
	user := m.User{}
	id, err := uuid.NewV4()
	db.Check(err)

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&user) // store details as user
	db.Check(err)

	// set up other user fields
	user.ID = id.String() // convert UUID to string
	user.DateCreated = time.Now()
	user.DateEdited = time.Now()
	user.Status = "Active"

	fmt.Println("time: ", user.DateCreated) // testing - print user details

	fmt.Println(user) // testing - print user details
	json.NewEncoder(w).Encode(&user)
}
