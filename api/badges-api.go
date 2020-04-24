package api

import (
	"FYP_Backend/db"
	q "FYP_Backend/db/queries"
	m "FYP_Backend/model"

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
	var reward m.Reward
	userID := r.Header.Get("user_id")

	err := json.NewDecoder(r.Body).Decode(&reward)
	db.Check(err)
	res := "Error: Update Score Failed"
	switch reward.RewardType {
	case "addCornellNote":
		res = q.UpdateUserStatsAddCornellNote(reward.Score, userID)

	case "addNote":
		res = q.UpdateUserStatsAddNote(reward.Score, userID)

	case "addCue":
		res = q.UpdateUserStatsAddNote(reward.Score, userID)

	case "completedReview":
		res = q.UpdateUserStatsCompleteReview(reward.Score, reward.Added, userID)

	case "shareNote":

	case "shareCornellNote":
	default:

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
