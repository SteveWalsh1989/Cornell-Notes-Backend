package queries

import (
	"FYP_Proto_Backend/db"
	m "FYP_Proto_Backend/model"
	"context"
	"fmt"
	"time"
)

//Ctx ... context for db queries
var Ctx context.Context

//GetFolder ... using folder ID returns folder from db as Folder struct
func GetFolder(id string) m.Folder {
	conn := db.CreateConn()
	var folder m.Folder

	query := "SELECT * FROM Folders WHERE id=" + id
	rows, err := conn.Query(query)
	db.Check(err)
	for rows.Next() {
		if err := rows.Scan(&folder.Title, &folder.ID,
			&folder.DateCreated, &folder.DateEdited); err != nil {
			fmt.Println("Error", err)
		}
		//fmt.Println("folder: ", folder.Name)
	}
	err = rows.Err()
	db.Check(err)

	db.CloseConn(conn)
	return folder
}

//GetFoldersItems ... using folder ID returns folder from db as Folder struct
func GetFoldersItems(userID string) []m.FolderItem {

	conn := db.CreateConn()
	var folderItem m.FolderItem
	var folderItems []m.FolderItem

	// Build Query
	query := "SELECT f.id AS folder_id, f.title AS folder_title, fi.item_type, fi.item_id " +
		"FROM folders f JOIN folder_users fu " +
		"ON (fu.folder_id = f.id) " +
		"JOIN folder_items fi " +
		"ON f.id = fi.folder_id " +
		"WHERE fu.user_id = '" + userID + "'"
	// fmt.Println("query: ", query)
	// Run Query
	rows, err := conn.Query(query)
	db.Check(err)
	// Assemble Results
	for rows.Next() {
		if err := rows.Scan(&folderItem.ID, &folderItem.Title,
			&folderItem.ItemType, &folderItem.ItemID); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			fmt.Println("Error: ", err)
		}
		folderItems = append(folderItems, folderItem)
	}
	err = rows.Err()
	db.Check(err)
	db.CloseConn(conn)
	// Return Results
	return folderItems
}

//UpdateFolderName ... Updates name of folder
func UpdateFolderName(id string, name string) {
	conn := db.CreateConn()

	stmt, err := conn.Prepare("UPDATE folders SET Title = ? WHERE id = ?")
	db.Check(err)
	_, err = stmt.Exec(name, id)
	db.Check(err)

	db.CloseConn(conn)

}

//DeleteFolder ... Updates folder status for 'deleted'
func DeleteFolder(folderID string) {
	conn := db.CreateConn()
	time := time.Now() // current time for time_edited
	_, err := conn.Query("UPDATE folders f SET f.status = 'Deleted', f.date_edited = ? WHERE f.id = ?", time, folderID)
	db.Check(err)
	db.CloseConn(conn)

}
