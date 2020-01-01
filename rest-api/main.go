package main

import (
	"FYP_Proto_Backend/api"
	m "FYP_Proto_Backend/model"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// mock data for testing
	fmt.Println("Checking works")
	m.Folders = append(m.Folders, m.Folder{ID: "1", Name: "Secuity Notes", Notes: "Testing 124"})
	m.Folders = append(m.Folders, m.Folder{ID: "2", Name: "Computer Architecture", Notes: "TTesting 456"})

	router := mux.NewRouter()

	/*
	 * FOLDERS
	 */
	router.HandleFunc("/folders", api.GetFolders).Methods("GET")
	router.HandleFunc("/folder", api.CreateFolder).Methods("POST")
	router.HandleFunc("/folders/{id}", api.GetFolder).Methods("GET")
	router.HandleFunc("/folders/{id}", api.UpdateFolderName).Methods("PUT")
	router.HandleFunc("/folders/{id}", api.DeleteFolder).Methods("DELETE")

	// Run server
	http.ListenAndServe(":8000", router)
}
