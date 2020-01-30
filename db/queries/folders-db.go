package queries

import (
	"FYP_Proto_Backend/db"
	m "FYP_Proto_Backend/model"
	"context"
	"fmt"
	"time"
)

var ctx context.Context

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

//UpdateFolderName ... Updates name of folder
func UpdateFolderName(name string) {
	conn := db.CreateConn()

	// test
	userID := 123332

	result, err := conn.ExecContext(ctx, "UPDATE folder SET name = ?  WHERE user_id = ?", name, userID)
	db.Check(err)

	rows, err := result.RowsAffected()
	db.Check(err)

	if rows != 1 {
		fmt.Printf("expected to affect 1 row, affected %d \n", rows)
	}

	db.CloseConn(conn)

}

//DeleteFolder ... Updates folder status for 'deleted'
func DeleteFolder(folderID string) {
	conn := db.CreateConn()

	time := time.Now() // current time for time_edited
	_, err := conn.Exec("UPDATE FOLDER SET status='Deleted, date_edited=$2  WHERE id=$1", folderID, time)
	db.Check(err)
	db.CloseConn(conn)

}
