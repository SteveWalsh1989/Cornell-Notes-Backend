package queries

import (
	"FYP_Backend/db"
	m "FYP_Backend/model"
	"fmt"
	"time"
)

//CreateCornellNote ... creates new cornell note for a user
func CreateCornellNote(cornellNote m.CornellNote, userID string, folderID string) string {
	fmt.Println(" -- CreateCornellNote DB")

	conn := db.CreateConn()
	tx, err := conn.Begin()
	db.Check(err)
	// 1: update cornell_note
	stmt, err := tx.Prepare("INSERT INTO cornell_notes (id, title, date_created, date_edited) VALUES(?,?,?,?);")
	if err != nil {
		fmt.Println("OOps - CreateCornellNote-  preparing statement 1 ", err)
		tx.Rollback()
		return "Error"
	}
	defer stmt.Close()
	if _, err := stmt.Exec(cornellNote.ID, cornellNote.Title, cornellNote.DateCreated, cornellNote.DateEdited); err != nil {
		fmt.Println("OOps - CreateCornellNote-  executing statement 1 ", err)
		tx.Rollback()
		return "Error"
	}
	// 2 : update cornell_users
	stmt, err = tx.Prepare("INSERT INTO cornell_users (cornell_note_id, user_id) VALUES(?,?);")
	if err != nil {
		fmt.Println("OOps - CreateCornellNote-  preparing statement 2 ", err)
		tx.Rollback()
		return "Error"
	}
	defer stmt.Close()
	if _, err := stmt.Exec(cornellNote.ID, userID); err != nil {
		fmt.Println("OOps - CreateCornellNote-  executing statement 2 ", err)
		tx.Rollback()
		return "Error"
	}
	// 3 : update folder_items
	stmt, err = tx.Prepare("INSERT INTO folder_items (folder_id, item_id, item_type) VALUES(?,?,?);")
	if err != nil {
		fmt.Println("OOps - CreateCornellNote-  preparing statement 3 ", err)
		tx.Rollback()
		return "Error"
	}
	defer stmt.Close()
	if _, err := stmt.Exec(folderID, cornellNote.ID, "CornellNote"); err != nil {
		fmt.Println("OOps - CreateCornellNote-  executing statement 3 ", err)
		tx.Rollback()
		return "Error"
	}

	tx.Commit()

	return "Note Added"
}

//GetCornellNoteSummary ... gets summary of note using ID
func GetCornellNoteSummary(cornellNoteID string, userID string) string {

	conn := db.CreateConn()
	summary := ""
	query := "SELECT cn.summary FROM cornell_notes cn JOIN cornell_users cnu ON cn.id = cnu.cornell_note_id " +
		"WHERE cnu.user_id = '" + userID + "' AND cn.ID = '" + cornellNoteID + "'"
	// Run Query
	rows, err := conn.Query(query)
	db.Check(err)
	// Assemble Results
	for rows.Next() {
		if err := rows.Scan(&summary); err != nil {
			fmt.Println("Error: ", err)
		}
	}
	err = rows.Err()
	db.Check(err)
	db.CloseConn(conn)

	// Return Results
	return summary

}

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

// GetCornellNoteTags ... returns tags per note
func GetCornellNoteTags(noteID string, userID string) []m.Tag {
	var tags []m.Tag
	var tag m.Tag
	conn := db.CreateConn()
	// Build Query
	query := "SELECT t.id, t.title, t.color " +
		"FROM users u JOIN tags t  ON u.id = t.user_id " +
		"JOIN tag_items ti on t.id = ti.tag_id " +
		"WHERE ti.item_id = '" + noteID + "' AND u.id = '" + userID + "'"
	// Run Query
	rows, err := conn.Query(query)
	db.Check(err)
	// Assemble Results
	for rows.Next() {
		if err := rows.Scan(&tag.ID, &tag.Title, &tag.Color); err != nil {
			fmt.Println("Error: ", err)
		}
		tags = append(tags, tag)
	}
	err = rows.Err()
	db.Check(err)
	db.CloseConn(conn)

	return tags
}

// UpdateCornellNote ... updates a cornell note details
func UpdateCornellNote(cornellNoteDetails m.UpdateCornellNote, userID string) string {
	res := ""

	conn := db.CreateConn()
	tx, err := conn.Begin()
	db.Check(err)
	stmt, err := tx.Prepare("UPDATE folder_items fi INNER JOIN folders f ON fi.folder_id = f.id SET fi.folder_id = ?   WHERE fi.item_id = ?;	")
	if err != nil {
		fmt.Println("OOps 1 preparing statement", err)
		tx.Rollback()
		return "Error"
	}
	defer stmt.Close()
	if _, err := stmt.Exec(cornellNoteDetails.Folder.ID, cornellNoteDetails.ID); err != nil {
		fmt.Println("OOps 1 executing statement", err)
		tx.Rollback()
		return "Error"
	}
	stmt, err = tx.Prepare("UPDATE cornell_notes cn SET cn.title = ? WHERE cn.id = ?	")

	if err != nil {
		fmt.Println("OOps 2 preparing statement", err)
		tx.Rollback()
		return "Error"
	}
	defer stmt.Close()
	if _, err := stmt.Exec(cornellNoteDetails.Title, cornellNoteDetails.ID); err != nil {
		fmt.Println("OOps 2 executing statement", err)
		tx.Rollback()
		return "Error"
	}

	tx.Commit()
	res = "Note Updated"
	return res
}

// DeleteCornellnote ... deletes cornell note using id - soft delete so sets status to deleted
func DeleteCornellnote(noteID string, userID string) bool {
	fmt.Println(" -- DeleteCornellNote DB")
	noteDeleted := false

	time := time.Now() // current time for time_edited

	// create transaction to update the item status for the cornell note and for folder items

	conn := db.CreateConn()
	tx, err := conn.Begin()
	db.Check(err)
	stmt, err := tx.Prepare("UPDATE cornell_notes cn SET cn.status = 'Deleted', cn.date_edited = ? WHERE cn.id = ?")
	if err != nil {
		fmt.Println("OOps preparing statement", err)
		tx.Rollback()
		return false
	}
	defer stmt.Close()
	if _, err := stmt.Exec(time, noteID); err != nil {
		fmt.Println("OOps executing statement", err)
		tx.Rollback()
		return false
	}
	stmt, err = tx.Prepare("UPDATE folder_items fi SET fi.status = 'Deleted' WHERE fi.item_id = ?")
	if err != nil {
		fmt.Println("OOps preparing statement", err)
		tx.Rollback()
		return false
	}
	defer stmt.Close()
	if _, err := stmt.Exec(noteID); err != nil {
		fmt.Println("OOps executing statement", err)
		tx.Rollback()
		return false
	}

	tx.Commit()
	noteDeleted = true // set true if no errors are returned

	return noteDeleted
}
