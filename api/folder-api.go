package api

import (
	q "FYP_Backend/db/queries"
	m "FYP_Backend/model"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetFolders : returns all folders
func GetFolders(w http.ResponseWriter, r *http.Request) {
	var folders []m.Folder
	var folderItem m.FolderItem
	var folderItems []m.FolderItem
	var itemNames []m.FolderItem
	id, _ := r.URL.Query()["id"]
	// Get folders itemIDs and item_types
	folderItems = q.GetFoldersItems(id[0])
	//fmt.Println("GetFolders: About to call q.GetNoteTitle") // testing
	itemNames = append(itemNames, q.GetNoteTitle(id[0])...)
	itemNames = append(itemNames, q.GetCornellNoteTitle(id[0])...)
	// Appends the item name to overall folder item
	for i, item := range folderItems {
		for _, name := range itemNames {
			if item.ItemID == name.ID {
				folderItems[i].ItemTitle = name.Title // add name to folder item
			}
		}
	}
	for _, item := range folderItems {
		exists := false
		for i, folder := range folders {
			if folder.ID == item.ID {
				exists = true // add to existing folder
				folderItem.Title = item.Title
				folderItem.ID = item.ID
				folderItem.ItemTitle = item.ItemTitle
				folderItem.ItemID = item.ItemID
				folderItem.ItemType = item.ItemType
				folders[i].Items = append(folders[i].Items, folderItem)
			}
		}
		if !exists { // create new folder folder
			var newfolder m.Folder
			folderItem.Title = item.Title
			folderItem.ID = item.ID
			newfolder.ID = item.ID
			newfolder.Title = item.Title
			folderItem.ItemTitle = item.ItemTitle
			folderItem.ItemID = item.ItemID
			folderItem.ItemType = item.ItemType
			newfolder.Items = append(newfolder.Items, folderItem)
			folders = append(folders, newfolder)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folders)
}

// GetFolder : returns folder by id
func GetFolder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var folder m.Folder
	folder = q.GetFolder(params["ID"])

	fmt.Println("GetFolder: folder name: ", folder.Title)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folder)
}

// CreateFolder : creates and returns folder
func CreateFolder(w http.ResponseWriter, r *http.Request) {

	var folder m.Folder
	_ = json.NewDecoder(r.Body).Decode(folder)

	folder.ID = strconv.Itoa(rand.Intn(1000000))
	//folder.Name = r.Body.

	m.Folders = append(m.Folders, folder)
	json.NewEncoder(w).Encode(&folder)
}

// UpdateFolderName : creates and returns folder
func UpdateFolderName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Check params arent empty
	if (params["id"] != "") && (params["name"] != "") {

		w.WriteHeader(422) // set HTTP response to bad entity
		json.NewEncoder(w).Encode("Missing Parameters")
		return
	}

	q.UpdateFolderName(params["id"], params["name"])

	json.NewEncoder(w).Encode(params["name"])
}

// DeleteFolder : creates and returns folder
func DeleteFolder(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	fmt.Println("Reached")

	fmt.Println("ID: ", params)

	q.DeleteFolder(params["id"])

	json.NewEncoder(w).Encode(m.Folders)
}
