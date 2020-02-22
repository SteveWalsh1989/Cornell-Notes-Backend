package queries

import (
	"FYP_Proto_Backend/db"
	m "FYP_Proto_Backend/model"
	"fmt"
)

//GetNoteName ...gets name of note using ID
func GetPDFsTitle(ID string) []m.FolderItem {
	conn := db.CreateConn()
	var item m.FolderItem
	var items []m.FolderItem
	// Build Query
	query := "SELECT p.id, .title FROM pdfs p JOIN pdf_users pu  ON p.id = pu.pdf_id WHERE pt.user_id = '" + ID + "'"
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
