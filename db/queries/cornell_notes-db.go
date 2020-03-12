package queries

import (
	"FYP_Backend/db"
	m "FYP_Backend/model"
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

// GetCornellNoteCues ... returns cornell notes cues and anwers
func GetCornellNoteCues(noteID string, userID string) m.CornellNote {
	var cornellNote m.CornellNote
	var cornellCues []m.CornellCue
	var cornellCue m.CornellCue
	conn := db.CreateConn()

	// Build Query
	query := "SELECT cn.title, cc.id AS CueID, cc.cue, cc.answer from sys.cornell_notes cn " +
		"JOIN sys.cornell_cues cc ON cn.id = cc.cornell_note_id " +
		"JOIN sys.cornell_users cu  ON cc.cornell_note_id = cu.cornell_note_id " +
		"JOIN sys.users u ON cu.user_id = u.id " +
		"WHERE  cn.id = '" + noteID + "' AND u.id = '" + userID + "'"

	// Run Query
	rows, err := conn.Query(query)
	db.Check(err)
	// Assemble Results
	for rows.Next() {
		if err := rows.Scan(&cornellNote.Title, &cornellCue.ID, &cornellCue.Cue, &cornellCue.Answer); err != nil {
			fmt.Println("Error: ", err)
		}
		cornellCues = append(cornellCues, cornellCue)
	}
	err = rows.Err()
	db.Check(err)
	db.CloseConn(conn)

	// add cues to note
	cornellNote.ID = noteID
	cornellNote.Cues = cornellCues

	return cornellNote
}

// GetCornellNoteTags ... returns tags per note
func GetCornellNoteTags(noteID string, userID string) []m.Tag {
	var tags []m.Tag
	var tag m.Tag
	conn := db.CreateConn()
	// Build Query
	query := "SELECT t.title, t.id, t.color " +
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

// UpdateCornellNote ... updates a cornell note
func UpdateCornellNote(noteID string, userID string) string {
	res := ""

	return res
}

// UpdateCornellNoteCue ... updates a cornell note cue
func UpdateCornellNoteCue(cue m.CornellCue, userID string) string {
	fmt.Println("test date again - ", cue.DateEdited)
	res := ""
	conn := db.CreateConn()
	tx, err := conn.Begin()
	db.Check(err)
	stmt, err := tx.Prepare("UPDATE cornell_cues cc " +
		"JOIN sys.cornell_users cu ON cc.cornell_note_id = cu.cornell_note_id " +
		"SET cc.cue = ?, cc.answer = ?  " +
		"WHERE cu.user_id = ? AND cc.id = ?;")
	if err != nil {
		fmt.Println("OOps2", err)
		tx.Rollback()
		return "Error"
	}
	defer stmt.Close()
	if _, err := stmt.Exec(cue.Cue, cue.Answer, userID, cue.ID); err != nil {
		fmt.Println("OOps3", err)

		tx.Rollback() // return an error too, we may want to wrap them
		return "Error"
	}
	stmt, err = tx.Prepare("UPDATE cornell_notes cn JOIN cornell_users cu ON cn.id = cu.cornell_note_id " +
		"SET cn.date_edited = ? WHERE cn.id = ? AND cu.user_id = ?")
	if err != nil {
		fmt.Println("OOps4", err)

		tx.Rollback()
		return "Error"
	}
	defer stmt.Close()

	if _, err := stmt.Exec(cue.DateEdited, cue.CornellNoteID, userID); err != nil {
		fmt.Println("OOps5", err)

		tx.Rollback() // return an error too, we may want to wrap them
		return "Error"
	}
	tx.Commit()

	return res
}
