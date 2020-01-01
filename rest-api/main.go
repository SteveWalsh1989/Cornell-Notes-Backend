package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Sample struct for folder to test router
type Folder struct {
	ID string `json:"id"`
	Name string `json:"title"`
	ParentFolder string `json:"body"`
	ChildFolder string `json:"body"`
  }
  var folders []Folder


func main() {

	fmt.Println("Checking works")
	router := mux.NewRouter()


	/*
	 * FOLDERS
	 */
	router.HandleFunc("/folders", getFolders).Methods("GET")
	router.HandleFunc("/folder",  createFolder).Methods("POST")
	router.HandleFunc("/folders/notes", getFolderNotes).Methods("GET")
	router.HandleFunc("/folders/{id}", updateFolderName).Methods("PUT")
	router.HandleFunc("/folders/{id}", deleteFolder).Methods("DELETE")
	


	// Run server
	http.ListenAndServe(":8000", router)
}
