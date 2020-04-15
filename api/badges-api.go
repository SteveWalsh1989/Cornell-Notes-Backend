package api

import (
	"encoding/json"
	"net/http"
)

// GetUserStats ... returns all user stats
func GetUserStats(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("user_id")

	res := q.GetUserStats(userID)

	res = "ok"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// UpdateUserStats ... updates user stats
func UpdateUserStats(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("user_id")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folders)
}
