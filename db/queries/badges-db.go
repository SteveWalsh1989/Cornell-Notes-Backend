package queries

import (
	"FYP_Backend/db"
	m "FYP_Backend/model"
	"fmt"
)

// GetBadges .. gets all badges
func GetBadges() []m.Badge {
	fmt.Println("Reached db query")

	var badge m.Badge
	var badges []m.Badge
	conn := db.CreateConn()

	query := "SELECT * FROM badges"

	rows, err := conn.Query(query)
	db.Check(err)
	for rows.Next() {
		if err := rows.Scan(&badge.ID, &badge.Title, &badge.Requirement); err != nil {
			fmt.Println("Error", err)
		}
		badges = append(badges, badge)
	}
	err = rows.Err()
	db.Check(err)
	db.CloseConn(conn)
	fmt.Println("badges", badges)

	return badges

}

// GetUserStats ... retrives stats per userID
func GetUserStats(userID string) m.UserStats {
	var userStats m.UserStats
	conn := db.CreateConn()

	query := "SELECT * FROM user_scores us WHERE us.user_id= '" + userID + "'"

	rows, err := conn.Query(query)
	db.Check(err)
	for rows.Next() {
		if err := rows.Scan(&userStats.UserID, &userStats.Points, &userStats.NotesCreated,
			&userStats.CornellNotesCreated, &userStats.NotesShared, &userStats.CornellNotesShared,
			&userStats.CuesCreated, &userStats.ReviewsCompleted, &userStats.CuesReviewed); err != nil {
			fmt.Println("Error", err)
		}
	}
	err = rows.Err()
	db.Check(err)
	db.CloseConn(conn)

	return userStats

}
