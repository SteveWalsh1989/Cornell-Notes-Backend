package api

import (
	q "FYP_Backend/db/queries"
	"encoding/json"
	"net/http"
)

// GetBadges ... gets badges
func GetBadges(w http.ResponseWriter, r *http.Request) {

	badges := q.GetBadges()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(badges)
}

// GetUserStats ... returns all user stats
func GetUserStats(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("user_id")

	stats := q.GetUserStats(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

// UpdateUserStats ... updates user stats
func UpdateUserStats(w http.ResponseWriter, r *http.Request) {

	// userID := r.Header.Get("user_id")

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(folders)
}
