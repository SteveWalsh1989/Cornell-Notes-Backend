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

// GetCornellNote -  GetCornell Note cues and answers
func GetCornellNote(w http.ResponseWriter, r *http.Request) {
	var cornellNote m.CornellNote
	var tags []m.Tag
	// Get query parameters
	noteID, _ := r.URL.Query()["note_id"]
	userID, _ := r.URL.Query()["user_id"]
	cornellNote = q.GetCornellNoteCues(noteID[0], userID[0])
	tags = q.GetCornellNoteTags(noteID[0], userID[0])
	cornellNote.Tags = tags
	// fmt.Println("GetTags: ", tags) // testing
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cornellNote)
}

// CreateCornellNote - create new Cornell Note
func CreateCornellNote(w http.ResponseWriter, r *http.Request) {
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

// UpdateCornellNote - update cornell note
func UpdateCornellNote(w http.ResponseWriter, r *http.Request) {
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

// DeleteCornellNote - deletes cornell note using id
func DeleteCornellNote(w http.ResponseWriter, r *http.Request) {
	ID, _ := r.URL.Query()["tagId"]
	userID, _ := r.URL.Query()["userId"]

	var tag m.Tag
	tag.ID = ID[0]
	tag.DateEdited = time.Now()
	tag.Status = "Deleted"
	deleted := q.DeleteTag(tag, userID[0])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deleted)
}
