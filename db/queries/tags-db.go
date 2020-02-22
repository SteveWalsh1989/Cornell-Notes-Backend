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

	query := "SELECT t.id FROM Tags t JOIN tag_users u ON t.id = WHERE id=" + id
	rows, err := conn.Query(query)
	db.Check(err)

	for rows.Next() {
		if err := rows.Scan(&tag.ID, &tag.Title); err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println("tag: ", tag.Title)
		tags = append(tags, tag) // add tag to list of tags
	}
	err = rows.Err()
	db.Check(err)

	db.CloseConn(conn)
	return tags
}

//CreateTag ... add new tag
func CreateTag(tag m.Tag, user_id string) m.Tag {

	return tag
}

//UpdateTag ... update tag name
func UpdateTag(tag m.Tag) m.Tag {

	return tag
}

//AddTagItem ... add tag to new item
func AddTagItem(tagID string, itemID string) {

}

//DeleteTag ... deletes tag by id
func DeleteTag(tagID string) {

}
