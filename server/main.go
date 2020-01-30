package main

import (
	"FYP_Proto_Backend/api"
	"FYP_Proto_Backend/db"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

// CORSRouterDecorator applies CORS headers to mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

var port = ":8011" //  port number for server

func main() {

	db.SetupDB()              // Initialise database
	router := mux.NewRouter() // create router

	/*
	 * FOLDERS
	 */
	router.HandleFunc("/folders", api.GetFolders).Methods("GET", "OPTIONS")
	router.HandleFunc("/folder", api.CreateFolder).Methods("POST", "OPTIONS")
	router.HandleFunc("/folders/", api.GetFolders).Methods("GET", "OPTIONS")
	router.HandleFunc("/folder/{id}", api.GetFolder).Methods("GET", "OPTIONS")
	router.HandleFunc("/folders/{id}", api.UpdateFolderName).Methods("PUT", "OPTIONS")
	router.HandleFunc("/folders/{id}", api.DeleteFolder).Methods("DELETE", "OPTIONS")

	// Run server
	db.LogValue("Started Server on port", port)
	http.ListenAndServeTLS(port, "localhost.crt", "localhost.key", &CORSRouterDecorator{router})
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
	db.LogRequest(req.Method, req.URL.Path)

	c.R.ServeHTTP(rw, req)
}
