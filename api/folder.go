package api

import (
	m "FYP_Proto_Backend/model"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetFolders : returns all folders
func GetFolders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m.Folders)
}

// GetFolder : returns folder by id
func GetFolder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range m.Folders {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
		return
	}
	json.NewEncoder(w).Encode(&m.Folder{})
}

// CreateFolder : creates and returns folder
func CreateFolder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var folder m.Folder
	_ = json.NewDecoder(r.Body).Decode(folder)
	folder.ID = strconv.Itoa(rand.Intn(1000000))
	m.Folders = append(m.Folders, folder)
	json.NewEncoder(w).Encode(&folder)
}

// UpdateFolderName : creates and returns folder
func UpdateFolderName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range m.Folders {
		if item.ID == params["id"] {
			m.Folders = append(m.Folders[:index], m.Folders[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(m.Folders)
}
