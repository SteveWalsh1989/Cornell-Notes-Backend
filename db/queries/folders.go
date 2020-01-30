package queries

import (
	"FYP_Proto_Backend/db"
	m "FYP_Proto_Backend/model"
	"fmt"
)

//GetFolder ... using folder ID returns folder from db as Folder struct
func GetFolder(id string) m.Folder {
	conn := db.CreateConn()
	var folder m.Folder

	query := "SELECT * FROM Folders WHERE id=" + id
	rows, err := conn.Query(query)
	db.Check(err)

	for rows.Next() {
		if err := rows.Scan(&folder.Name, &folder.ID, &folder.Status,
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

//GetFolders ... using folder ID returns folder from db as Folder struct
func GetFolders() []m.Folder {
	conn := db.CreateConn()
	var folder m.Folder
	var folders = []m.Folder{}

	rows, err := conn.Query("SELECT * FROM Folders")
	db.Check(err)

	for rows.Next() {
		if err := rows.Scan(&folder.ID, &folder.Name, &folder.Status,
			&folder.DateCreated, &folder.DateEdited); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			fmt.Println("Error: ", err)
		}
		// fmt.Println("folder: ", folder.Name)
		folders = append(folders, folder)
	}
	err = rows.Err()
	db.Check(err)

	db.CloseConn(conn)
	return folders
}
