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
	router.HandleFunc("/updateFolderName/{id}/{name}", api.UpdateFolderName).Methods("PUT", "OPTIONS")
	router.HandleFunc("/folder/{id}", api.DeleteFolder).Methods("DELETE", "OPTIONS")
}
