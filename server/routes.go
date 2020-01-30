package main

import (
	"FYP_Proto_Backend/api"

	"github.com/gorilla/mux"
)

// ServeRoutes .. handles API Routes  
func ServeRoutes(router *mux.Router) {

	/*
	 * FOLDERS
	 */
	router.HandleFunc("/folders", api.GetFolders).Methods("GET", "OPTIONS")
	router.HandleFunc("/folder", api.CreateFolder).Methods("POST", "OPTIONS")
	router.HandleFunc("/folders/", api.GetFolders).Methods("GET", "OPTIONS")
	router.HandleFunc("/folder/{id}", api.GetFolder).Methods("GET", "OPTIONS")
	router.HandleFunc("/folders/{id}", api.UpdateFolderName).Methods("PUT", "OPTIONS")
	router.HandleFunc("/folders/{id}", api.DeleteFolder).Methods("DELETE", "OPTIONS")
}
