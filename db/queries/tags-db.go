package queries

import (
	"FYP_Proto_Backend/db"
	m "FYP_Proto_Backend/model"
	"fmt"
)

//GetTags ... using folder ID returns tag from
func GetTags(id string) []m.Tag {
	conn := db.CreateConn()
	var tag m.Tag
	var tags []m.Tag
	query := "SELECT t.id, t.title, t.color FROM tags t WHERE t.status = 'Active' AND t.user_id =" + id
	rows, err := conn.Query(query)
	db.Check(err)
	for rows.Next() {
		if err := rows.Scan(&tag.ID, &tag.Title, &tag.Color); err != nil {
			fmt.Println("Error", err)
		}
		tags = append(tags, tag) // add tag to list of tags
	}
	err = rows.Err()
	db.Check(err)

	db.CloseConn(conn)
	return tags
}

//CreateTag ... add new tag
func CreateTag(tag m.Tag, userID string) m.Tag {
	conn := db.CreateConn()
	stmt, err := conn.Prepare("INSERT INTO tags (id, title, user_id, color, date_created, date_edited) VALUES(?,?,?,?,?,?);")
	db.Check(err)
	_, errr := stmt.Exec(tag.ID, tag.Title, userID, tag.Color, tag.DateCreated, tag.DateEdited)
	db.Check(errr)
	db.CloseConn(conn)
	return tag
}

//UpdateTag ... update tag name
func UpdateTag(tag m.Tag, userID string) m.Tag {
	conn := db.CreateConn()
	stmt, err := conn.Prepare("UPDATE tags SET title=?, color=?, date_edited=? WHERE id=? AND user_id=?;")
	db.Check(err)
	_, errr := stmt.Exec(tag.Title, tag.Color, tag.DateEdited, tag.ID, userID)
	db.Check(errr)
	db.CloseConn(conn)
	return tag
}

//AddTagItem ... add tag to new item
func AddTagItem(tagID string, itemID string) {

}

//DeleteTag ... deletes tag by id
func DeleteTag(tag m.Tag, userID string) bool {
	conn := db.CreateConn()
	stmt, err := conn.Prepare("UPDATE tags SET status=?, date_edited=? WHERE id=? AND user_id=?;")
	db.Check(err)
	_, errr := stmt.Exec(tag.Status, tag.DateEdited, tag.ID, userID)
	db.Check(errr)
	db.CloseConn(conn)
	return true
}
