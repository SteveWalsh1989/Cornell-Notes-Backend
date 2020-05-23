package api

import (
	"FYP_Backend/db"
	q "FYP_Backend/db/queries"
	m "FYP_Backend/model"
	"encoding/json"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

// GetFriendsList ... gets list of user friends
func GetFriendsList(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("user_id")

	friendsList := m.Users

	friendsList = q.GetFriendsList(userID)

	json.NewEncoder(w).Encode(friendsList)
}

// AddFriend ... add new friend
func AddFriend(w http.ResponseWriter, r *http.Request) {
	friend := m.User{}
	// userID := r.Header.Get("user_id")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&friend)
	db.Check(err)
	// check if there is an existing user by email and return error if true
	// set up other user fields
	id, err2 := uuid.NewV4() // create new UUID for new user
	db.Check(err2)
	friend.ID = id.String()
	friend.DateCreated = time.Now()
	friend.DateEdited = time.Now()
	//fmt.Println("time: ", user.DateCreated) // testing - print user details
	// q.CreateUser(friend, userID) // Call db query
	//fmt.Println(user) // testing - print user details
	json.NewEncoder(w).Encode(&friend)
}

// DeleteFriend ... remiove friends from users friends list
func DeleteFriend(w http.ResponseWriter, r *http.Request) {
	// userID := r.Header.Get("user_id")
	// // friendsList := m.Users
	// // friendsList = q.DeleteFriend(userID)
	// json.NewEncoder(w).Encode(friendsList)
}
