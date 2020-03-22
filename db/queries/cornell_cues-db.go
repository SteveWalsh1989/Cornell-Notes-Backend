package queries

import (
	"FYP_Backend/db"
	m "FYP_Backend/model"
	"fmt"
)

// GetCornellNoteCues ... returns cornell notes cues and anwers
func GetCornellNoteCues(noteID string, userID string) m.CornellNote {
	var cornellNote m.CornellNote
	var cornellCues []m.CornellCue
	var cornellCue m.CornellCue
	conn := db.CreateConn()

	// Build Query
	query := "SELECT cn.title, cc.id AS CueID, cc.cue, cc.cue_order, cc.answer from sys.cornell_notes cn " +
		"JOIN sys.cornell_cues cc ON cn.id = cc.cornell_note_id " +
		"JOIN sys.cornell_users cu  ON cc.cornell_note_id = cu.cornell_note_id " +
		"JOIN sys.users u ON cu.user_id = u.id " +
		"WHERE  cn.id = '" + noteID + "' AND u.id = '" + userID + "'" +
		"ORDER BY cc.cue_order ASC"

	// Run Query
	rows, err := conn.Query(query)
	db.Check(err)
	// Assemble Results
	for rows.Next() {
		if err := rows.Scan(&cornellNote.Title, &cornellCue.ID, &cornellCue.Cue, &cornellCue.CueOrder, &cornellCue.Answer); err != nil {
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

// AddCornellNoteCue ... adds a new a cornell note cue
func AddCornellNoteCue(cue m.CornellCue, userID string) m.CornellCue {
	fmt.Println("-- AddCornellNoteCue: ", cue)
	conn := db.CreateConn()
	stmt, err := conn.Prepare("INSERT INTO cornell_cues (id, cornell_note_id, cue_order, cue, answer) VALUES(?,?,?,?,?);")
	db.Check(err)
	_, errr := stmt.Exec(cue.ID, cue.CornellNoteID, cue.CueOrder, cue.Cue, cue.Answer)
	db.Check(errr)
	db.CloseConn(conn)
	return cue

}

// UpdateCornellNoteCue ... updates a cornell note cue
func UpdateCornellNoteCue(cue m.CornellCue, userID string) string {

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

// DeleteCornellNoteCue ... deletes a cue from cornell ote - hard delete
func DeleteCornellNoteCue(cue m.CornellCue, userID string) bool {
	deleted := false
	conn := db.CreateConn()

	stmt, err := conn.Prepare("DELETE FROM cornell_cues cc WHERE EXISTS( SELECT 1 FROM cornell_users cu Where cc.cornell_note_id = cu.cornell_note_id AND cc.id = ? AND cu.user_id = ?);")

	db.Check(err)
	_, errr := stmt.Exec(cue.ID, userID)
	db.Check(errr)
	db.CloseConn(conn)
	deleted = true
	return deleted
}
