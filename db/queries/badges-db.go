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
		if err := rows.Scan(&badge.ID, &badge.Title, &badge.Requirement, &badge.Icon); err != nil {
			fmt.Println("Error", err)
		}
		badges = append(badges, badge)
	}
	err = rows.Err()
	db.Check(err)
	db.CloseConn(conn)

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

// UpdateUserStatsAddCornellNote ... updates stats per userID
func UpdateUserStatsAddCornellNote(points int, userID string) string {

	conn := db.CreateConn()

	stmt, err := conn.Prepare("UPDATE user_scores us SET us.cornell_notes_created = us.cornell_notes_created + 1 , us.points = us.points + ? WHERE us.user_id= ?;")

	db.Check(err)
	_, errr := stmt.Exec(points, userID)
	db.Check(errr)
	db.CloseConn(conn)

	return "Updated Score"

}

// UpdateUserStatsAddNote ... updates stats per userID
func UpdateUserStatsAddNote(points int, userID string) string {

	conn := db.CreateConn()
	stmt, err := conn.Prepare("UPDATE user_scores us SET us.notes_created = us.notes_created + 1 , us.points = us.points + ? WHERE us.user_id= ?;")

	db.Check(err)
	_, errr := stmt.Exec(points, userID)
	db.Check(errr)
	db.CloseConn(conn)

	return "Updated Score"

}

// UpdateUserStatsCompleteReview ... updates stats per userID
func UpdateUserStatsCompleteReview(points int, userID string) string {
	conn := db.CreateConn()
	stmt, err := conn.Prepare("UPDATE user_scores us SET us.reviews_completed  = us.reviews_completed + 1 , us.points = points + ? WHERE us.user_id= ?;")

	fmt.Println("Statement:", stmt)

	db.Check(err)
	_, errr := stmt.Exec(points, userID)
	db.Check(errr)
	db.CloseConn(conn)

	return "Updated Score"
}
