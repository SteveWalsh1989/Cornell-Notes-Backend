package api

import (
	"FYP_Proto_Backend/db"
	q "FYP_Proto_Backend/db/queries"
	m "FYP_Proto_Backend/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

	var tag m.Tag
	// Build new tag object using passed in name
	id, err := uuid.NewV4() // create new UUID for new user
	db.Check(err)
	tag.ID = id.String()
	tag.Title = title[0]
	tag.Color = color[0]
	tag.DateCreated = time.Now()
	tag.DateEdited = time.Now()
	fmt.Println("CreateTag: ", tag)

	tag = q.CreateTag(tag, userID[0])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tag)
}

// UpdateTag - update tag name
func UpdateTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var folder m.Folder
	folder = q.GetFolder(params["ID"])

	fmt.Println("GetFolder: folder name: ", folder.Name)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folder)
}

// DeleteTag - deletes tag using id
func DeleteTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var folder m.Folder
	folder = q.GetFolder(params["ID"])

	fmt.Println("GetFolder: folder name: ", folder.Name)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folder)
}
