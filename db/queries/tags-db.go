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
	fmt.Println("id: ", id)
	query := "SELECT t.id, t.title, t.color FROM tags t JOIN tag_items ti ON t.id = ti.tag_id WHERE t.user_id =" + id
	fmt.Println("query: ", query)
	rows, err := conn.Query(query)
	db.Check(err)
	for rows.Next() {
		if err := rows.Scan(&tag.ID, &tag.Title, &tag.Color); err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println("here 22")
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
