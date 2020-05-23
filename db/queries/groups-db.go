package queries

import (
	"FYP_Backend/db"
	m "FYP_Backend/model"
	"fmt"
)

//GetFriendsList ... gets a users friends list
func GetFriendsList(userID string) []m.User {

	conn := db.CreateConn()
	var friend m.User
	friendsList := m.Users

	query := "SELECT u.id, u.user_name,  u.email " +
		"FROM users u JOIN friends f " +
		"ON u.id = f.friend_id " +
		"WHERE f.user_id = '" + userID + "';"
	// Run Query
	rows, err := conn.Query(query)
	db.Check(err)
	// Assemble Results
	for rows.Next() {
		if err := rows.Scan(&friend.ID, &friend.Name,
			&friend.Email); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			fmt.Println("Error: ", err)
		}
		friendsList = append(friendsList, friend)
	}
	err = rows.Err()
	db.Check(err)
	db.CloseConn(conn)
	// Return Results
	return friendsList
}
