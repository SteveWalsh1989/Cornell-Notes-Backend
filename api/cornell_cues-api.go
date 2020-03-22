package api

import (
	"FYP_Backend/db"
	q "FYP_Backend/db/queries"
	m "FYP_Backend/model"
	"encoding/json"
	"net/http"
	"time"
)

// AddCornellNoteCue - Adds new cornell Cue
func AddCornellNoteCue(w http.ResponseWriter, r *http.Request) {

	var cue m.CornellCue
	// Get Cue details from body and user id from header
	userID := r.Header.Get("user_id")
	err := json.NewDecoder(r.Body).Decode(&cue)
	db.Check(err)
	// Update DB with new Cue information
	res := q.AddCornellNoteCue(cue, userID) // run db query

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

// UpdateCornellNoteCue - update cornell cue
func UpdateCornellNoteCue(w http.ResponseWriter, r *http.Request) {
	var cue m.CornellCue
	// Get Cue details from body and user id from header
	userID := r.Header.Get("user_id")
	err := json.NewDecoder(r.Body).Decode(&cue)
	db.Check(err)
	cue.DateEdited = time.Now()
	// Update DB with new Cue information
	res := q.UpdateCornellNoteCue(cue, userID) // run db query

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
