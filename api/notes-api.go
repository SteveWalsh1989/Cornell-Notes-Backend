package api

import (
	"FYP_Backend/db"
	q "FYP_Backend/db/queries"
	m "FYP_Backend/model"
	"time"

	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// GetNote ... gets notes contents
func GetNote(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("user_id")
	noteID, _ := r.URL.Query()["note_id"]

	note := q.GetNote(noteID[0], userID) // run db query
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// SaveNote - add new note
func SaveNote(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("user_ID")
	// set up New Note
	var note m.Note
	id, err := uuid.NewV4() // create new UUID for new user
	note.ID = id.String()
	title, _ := r.URL.Query()["note_title"]
	note.Title = title[0]
	folderID, _ := r.URL.Query()["folder_ID"]
	note.DateCreated = time.Now()
	db.Check(err)
	res := q.SaveNote(note, folderID[0], userID) // run db query

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// UpdateNote - update  note
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("user_ID")
	var note m.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	db.Check(err)

	q.UpdateNote(note, userID) // run db query

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// DeleteNote - delete  note
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	var note m.Note

	err := json.NewDecoder(r.Body).Decode(&note)
	db.Check(err)

	fmt.Println("\n\ntesting BODY: ", note.Title)
	// res := q.UpdateCornellNote(note[0], userID[0]) // run db query

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}
