package main

import (
	"FYP_Proto_Backend/api"
	"FYP_Proto_Backend/db"
	m "FYP_Proto_Backend/model"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

// CORSRouterDecorator applies CORS headers to mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

func main() {
	// Initialise database
	db.SetupDB()

	// mock data for testing
	m.Folders = append(m.Folders, m.Folder{ID: "11", Name: "Microservices", Status: "Active"})
	m.Folders = append(m.Folders, m.Folder{ID: "22", Name: "Computer Architecture", Status: "Active"})

	router := mux.NewRouter()

	/*
	 * FOLDERS
	 */
	router.HandleFunc("/folders", api.GetFolders).Methods("GET", "OPTIONS")
	router.HandleFunc("/folder", api.CreateFolder).Methods("POST", "OPTIONS")
	router.HandleFunc("/folder/", api.GetFolder).Methods("GET", "OPTIONS")
	router.HandleFunc("/folders/{id}", api.UpdateFolderName).Methods("PUT", "OPTIONS")
	router.HandleFunc("/folders/{id}", api.DeleteFolder).Methods("DELETE", "OPTIONS")

	// Run server
	port := ":8011"
	fmt.Println("Started Server on port", port)
	http.ListenAndServeTLS(":8011", "localhost.crt", "localhost.key", &CORSRouterDecorator{router})

}

// ServeHTTP:
// 1 wraps the HTTP server enabling CORS headers for preflight requests
// 2 Wraps the HTTP server header for response to front end for regular requests
func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, YourOwnHeader")
		rw.Header().Set("TESTINGGGGGG", "I'm here!!")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// TEST: log url and request type
	log.Println("--Log: ", req.Method, " Request:  ", req.URL.Path)
	c.R.ServeHTTP(rw, req)
}

// Checks for non nil errors and prints
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}