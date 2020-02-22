package api

import (
	q "FYP_Proto_Backend/db/queries"
	m "FYP_Proto_Backend/model"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetFolders : returns all folders
func GetFolders(w http.ResponseWriter, r *http.Request) {
	var folderItems []m.FolderItem
	var itemNames []m.FolderItem
	id, _ := r.URL.Query()["id"]
	// Get folders itemIDs and item_types
	folderItems = q.GetFoldersItems(id[0])

	fmt.Println("GetFolders: About to call q.GetNoteTitle") // testing

	itemNames = append(itemNames, q.GetNoteTitle(id[0])...)
	itemNames = append(itemNames, q.GetCornellNoteTitle(id[0])...)
	//Names = append(Names, q.GetPDFNoteName(PDFIDs))
	fmt.Println("itemNames: ", itemNames) // testing

	// fmt.Println("GetFolders: ", folders) // testing
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folderItems)
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
