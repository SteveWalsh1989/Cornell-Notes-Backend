package main

import (
	"FYP_Proto_Backend/api"

	"github.com/gorilla/mux"
)

// ServeRoutes .. handles API Routes
func ServeRoutes(router *mux.Router) {

	/*
	 * Login / Register User
	 */
	router.HandleFunc("/register", api.RegisterUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", api.LoginUser).Methods("GET", "OPTIONS")

	/*
	 * FOLDERS
	 */
	router.HandleFunc("/folders", api.GetFolders).Methods("GET", "OPTIONS")
	router.HandleFunc("/folder", api.CreateFolder).Methods("POST", "OPTIONS")
	router.HandleFunc("/folder", api.UpdateFolderName).Methods("PUT", "OPTIONS")
	router.HandleFunc("/folders/", api.GetFolders).Methods("GET", "OPTIONS")
	router.HandleFunc("/folder/{id}", api.DeleteFolder).Methods("DELETE", "OPTIONS")

	/*
	 * Tags
	 */
	router.HandleFunc("/tags", api.GetTags).Methods("GET", "OPTIONS")
	router.HandleFunc("/tag", api.CreateTag).Methods("POST", "OPTIONS")
	router.HandleFunc("/tag", api.UpdateTag).Methods("PUT", "OPTIONS")
	router.HandleFunc("/tag", api.DeleteTag).Methods("DELETE", "OPTIONS")

	/*
	 * Cornell Notes
	 */
	router.HandleFunc("/cornellnote", api.GetCornellNote).Methods("GET", "OPTIONS")
	router.HandleFunc("/cornellnote", api.CreateCornellNote).Methods("POST", "OPTIONS")
	router.HandleFunc("/cornellnote", api.UpdateCornellNote).Methods("PUT", "OPTIONS")
	router.HandleFunc("/cornellnote", api.DeleteCornellNote).Methods("DELETE", "OPTIONS")

	/*
	 * Cornell Note Cues
	 */
	//  router.HandleFunc("/cornellnote/cue", api.addCornellNoteCue).Methods("POST", "OPTIONS")
	router.HandleFunc("/cornellnote/cue", api.UpdateCornellNoteCue).Methods("PUT", "OPTIONS")
	//  router.HandleFunc("/cornellnote/cue", api.DeleteCornellNoteCue).Methods("DELETE", "OPTIONS")

}
