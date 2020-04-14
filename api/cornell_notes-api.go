package api

import (
	"FYP_Backend/db"
	q "FYP_Backend/db/queries"
	m "FYP_Backend/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GetCornellNote -  GetCornell Note cues and answers
func GetCornellNote(w http.ResponseWriter, r *http.Request) {

	var cornellNote m.CornellNote
	var tags []m.Tag
	// Get query parameters
	noteID, _ := r.URL.Query()["note_id"]
	userID, _ := r.URL.Query()["user_id"]
	cornellNote = q.GetCornellNoteCues(noteID[0], userID[0])
	summary := q.GetCornellNoteSummary(noteID[0], userID[0])
	tags = q.GetCornellNoteTags(noteID[0], userID[0])
	cornellNote.Tags = tags
	cornellNote.Summary = summary
	// fmt.Println("GetTags: ", tags) // testing

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cornellNote)
}

// CreateCornellNote - create new Cornell Note
func CreateCornellNote(w http.ResponseWriter, r *http.Request) {
	var cornellNote m.CornellNote

	// Get Cue details from body and user id from header

	userID := r.Header.Get("user_id")
	folderID := r.Header.Get("folder_id")

	_ = json.NewDecoder(r.Body).Decode(&cornellNote)
	cornellNote.DateCreated = time.Now()
	cornellNote.DateEdited = time.Now()
	fmt.Println(" -- folderID", folderID)

	// DB query
	res := q.CreateCornellNote(cornellNote, userID, folderID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// UpdateCornellNote - update cornell note
func UpdateCornellNote(w http.ResponseWriter, r *http.Request) {
	var cornellNote m.UpdateCornellNote
	// Get Cue details from body and user id from header
	userID := r.Header.Get("user_id")
	err := json.NewDecoder(r.Body).Decode(&cornellNote)
	db.Check(err)
	// DB query
	res := q.UpdateCornellNote(cornellNote, userID) // run db query

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// DeleteCornellNote - deletes cornell note using id
func DeleteCornellNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -- DeleteCornellNote API")
	noteID, _ := r.URL.Query()["note_id"]
	userID, _ := r.URL.Query()["user_id"]
	deleted := q.DeleteCornellnote(noteID[0], userID[0])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deleted)
}
