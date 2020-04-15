package queries

import (
	"FYP_Backend/db"
	m "FYP_Backend/model"
	"fmt"
)

// GetNote .. gets notes contents from db
func GetNote(noteID string, userID string) m.Note {
	var note m.Note
	conn := db.CreateConn()
	query := "SELECT n.title, n.id, n.body, n.date_created, n.date_edited " +
		"FROM notes n JOIN note_users nu on n.id = nu.note_id " +
		"WHERE nu.user_id ='" + userID + "' AND n.id='" + noteID + "'"
	rows, err := conn.Query(query)
	db.Check(err)
	for rows.Next() {
		if err := rows.Scan(&note.Title, &note.ID, &note.Body,
			&note.DateCreated, &note.DateEdited); err != nil {
			fmt.Println("Error", err)
		}
	}
	err = rows.Err()
	db.Check(err)

	db.CloseConn(conn)
	return note
}

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

// SaveNote ... adds new note to DB
func SaveNote(note m.Note, folderID string, userID string) string {
	res := ""
	conn := db.CreateConn()
	tx, err := conn.Begin()
	db.Check(err)
	stmt, err := tx.Prepare("INSERT INTO folder_items (folder_id, item_id, item_type) VALUES (?, ?, ?) ;")
	if err != nil {
		fmt.Println("OOps2", err)
		tx.Rollback()
		return "Error"
	}
	defer stmt.Close()
	if _, err := stmt.Exec(folderID, note.ID, "Note"); err != nil {
		fmt.Println("OOps3", err)

		tx.Rollback() // return an error too, we may want to wrap them
		return "Error"
	}
	stmt, err = tx.Prepare("INSERT INTO notes (id, title, date_created) VALUES (?, ?, ?) ;")
	if err != nil {
		fmt.Println("OOps4", err)
		tx.Rollback()
		return "Error"
	}
	defer stmt.Close()
	if _, err := stmt.Exec(note.ID, note.Title, note.DateCreated); err != nil {
		fmt.Println("OOps5", err)
		tx.Rollback() // return an error too, we may want to wrap them
		return "Error"
	}
	stmt, err = tx.Prepare("INSERT INTO note_users (note_id, user_id) VALUES (?, ?) ;")
	if err != nil {
		fmt.Println("OOps6", err)
		tx.Rollback()
		return "Error"
	}
	defer stmt.Close()
	if _, err := stmt.Exec(note.ID, userID); err != nil {
		fmt.Println("OOps7", err)
		tx.Rollback() // return an error too, we may want to wrap them
		return "Error"
	}
	tx.Commit()
	return res
}

// UpdateNote ... updates note
func UpdateNote(note m.Note, userID string) {

	conn := db.CreateConn()
	stmt, err := conn.Prepare("UPDATE notes n JOIN note_users nu ON n.id = nu.user_id SET n.title=?, n.body=?, n.date_edited=? WHERE n.id=? AND nu.user_id=?;")
	db.Check(err)
	_, errr := stmt.Exec(note.Title, note.Body, note.DateEdited, note.ID, userID)
	db.Check(errr)
	db.CloseConn(conn)

}
