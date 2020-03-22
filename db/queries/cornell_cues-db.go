package queries

import (
	"FYP_Backend/db"
	m "FYP_Backend/model"
	"fmt"
)

// AddCornellNoteCue ... adds a new a cornell note cue
func AddCornellNoteCue(cue m.CornellCue, userID string) m.CornellCue {
	fmt.Println("-- AddCornellNoteCue: ", cue)
	conn := db.CreateConn()
	stmt, err := conn.Prepare("INSERT INTO cornell_cues (id, cornell_note_id, cue, answer) VALUES(?,?,?,?);")
	db.Check(err)
	_, errr := stmt.Exec(cue.ID, cue.CornellNoteID, cue.Cue, cue.Answer)
	db.Check(errr)
	db.CloseConn(conn)
	return cue

}

// UpdateCornellNoteCue ... updates a cornell note cue
func UpdateCornellNoteCue(cue m.CornellCue, userID string) string {
	fmt.Println("test  UpdateCornellNoteCue - date edited - ", cue.DateEdited)
	fmt.Println("test  UpdateCornellNoteCue - date created - ", cue.DateCreated)

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
