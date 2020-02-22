package queries

import (
	"FYP_Proto_Backend/db"
	m "FYP_Proto_Backend/model"
	"fmt"
)

// GetNoteTitle ...gets name of note using ID
func GetNoteTitle(ID string) []m.FolderItem {

	conn := db.CreateConn()
	var item m.FolderItem
	var items []m.FolderItem
	// Build Query
	query := "SELECT n.id, n.title FROM notes n JOIN note_users nu ON n.id = nu.note_id WHERE nu.user_id = '" + ID + "' AND n.status = 'Active'"
	// Run Query
	rows, err := conn.Query(query)
	db.Check(err)
	// Assemble Results
	for rows.Next() {
		if err := rows.Scan(&item.ID, &item.Title); err != nil {
			fmt.Println("Error: ", err)
		}
		items = append(items, item)
	}
	err = rows.Err()
	db.Check(err)
	db.CloseConn(conn)
	// Return Results
	return items
}
