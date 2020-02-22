package queries

import (
	"FYP_Proto_Backend/db"
	m "FYP_Proto_Backend/model"
	"fmt"
)

//GetCornellNoteTitle ...gets name of note using ID
func GetCornellNoteTitle(ID string) []m.FolderItem {
	conn := db.CreateConn()
	var item m.FolderItem
	var items []m.FolderItem
	// Build Query
	query := "SELECT cn.id, cn.title FROM cornell_notes cn JOIN cornell_users cnu ON cn.id = cnu.cornell_note_id " +
		"WHERE cnu.user_id = '" + ID + "' AND cn.status = 'Active'"
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
