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

// GetTags -  get user tags
func GetTags(w http.ResponseWriter, r *http.Request) {
	var tags []m.Tag
	// Get query parameters
	id, _ := r.URL.Query()["id"]
	tags = q.GetTags(id[0])
	// fmt.Println("GetTags: ", tags) // testing
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}

// CreateTag - create new tag
func CreateTag(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	title, _ := r.URL.Query()["title"]
	color, _ := r.URL.Query()["color"]
	userID, _ := r.URL.Query()["userId"]

	// create new UUID
	id, err := uuid.NewV4() // create new UUID for new user
	db.Check(err)

	// Build new tag object using query params
	var tag m.Tag
	tag.ID = id.String()
	tag.Title = title[0]
	tag.Color = color[0]
	tag.DateCreated = time.Now()
	tag.DateEdited = time.Now()
	//fmt.Println("CreateTag: ", tag)

	tag = q.CreateTag(tag, userID[0]) // run db query
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tag)
}

// UpdateTag - update tag name
func UpdateTag(w http.ResponseWriter, r *http.Request) {
	id, _ := r.URL.Query()["tagId"]
	title, _ := r.URL.Query()["title"]
	color, _ := r.URL.Query()["color"]
	userID, _ := r.URL.Query()["userId"]

	// Build new tag object using query params
	var tag m.Tag
	tag.ID = id[0]
	tag.Title = title[0]
	tag.Color = color[0]
	tag.DateEdited = time.Now()

	tag = q.UpdateTag(tag, userID[0]) // run db query

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tag)
}

// DeleteTag - deletes tag using id
func DeleteTag(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.URL.Query()["userId"]
	title, _ := r.URL.Query()["title"]

	q.DeleteTag(title[0], userID[0])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userID)
}
