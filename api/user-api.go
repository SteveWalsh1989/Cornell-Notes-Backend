package api

import (
	"FYP_Proto_Backend/db"
	m "FYP_Proto_Backend/model"
	"encoding/json"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// RegisterUser : registers new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {

	// Check form data correct

	// create new UUID for new user
	user := m.User{}
	id, err := uuid.NewV4()
	db.Check(err)

	user.ID = id.String() // convert UUID to string

	db.LogValue("UserID: ", user.ID) // testing - print UUID

	json.NewEncoder(w).Encode(&user)
}
