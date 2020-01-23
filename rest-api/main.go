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
	m.Folders = append(m.Folders, m.Folder{ID: "1", Name: "Security Notes", Notes: "Testing 123 ABC"})
	m.Folders = append(m.Folders, m.Folder{ID: "2", Name: "Computer Architecture", Notes: "Testing 124"})

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
	port := ":8010"
	fmt.Println("Started Server on port", port)
	//http.ListenAndServe(port, router)
	http.ListenAndServeTLS(":8010", "https-server.crt", "https-server.key", router)

	//certmagic.HTTPS([]string{"localhost:8010"}, router)

}
