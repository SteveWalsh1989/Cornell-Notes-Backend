package api

import (
	q "FYP_Proto_Backend/db/queries"
	m "FYP_Proto_Backend/model"
	"fmt"

	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetFolders : returns all folders
func GetFolders(w http.ResponseWriter, r *http.Request) {
	var folders []m.Folder
	folders = q.GetFolders()

	// fmt.Println("GetFolders: ", folders) // testing

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folders)
}

// GetFolder : returns folder by id
func GetFolder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var folder m.Folder
	folder = q.GetFolder(params["ID"])

	fmt.Println("GetFolder: folder name: ", folder.Name)
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
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	params := mux.Vars(r)
	for index, item := range m.Folders {
		if item.ID == params["id"] {
			m.Folders = append(m.Folders[:index], m.Folders[index+1:]...)
			var folder m.Folder
			_ = json.NewDecoder(r.Body).Decode(folder)
			folder.ID = params["id"]
			m.Folders = append(m.Folders, folder)
			json.NewEncoder(w).Encode(&folder)
			return
		}
	}
	json.NewEncoder(w).Encode(m.Folders)
}

// DeleteFolder : creates and returns folder
func DeleteFolder(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	params := mux.Vars(r)
	for index, item := range m.Folders {
		if item.ID == params["id"] {
			m.Folders = append(m.Folders[:index], m.Folders[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(m.Folders)
}
