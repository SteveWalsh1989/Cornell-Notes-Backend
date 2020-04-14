package main

import (
	"FYP_Backend/db"
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

	router := mux.NewRouter() // create router
	ServeRoutes(router)       // handles API Route requests

	db.LogTitle("Setting Up FYP Backend")
	db.LogValue("Starting Server on port", port)
	http.ListenAndServeTLS(port, "localhost.crt", "localhost.key", &CORSRouterDecorator{router}) // Run server
}

// ServeHTTP:
// 1 wraps the HTTP server enabling CORS headers for preflight requests
// 2 Wraps the HTTP server header for response to front end for regular requests
func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Add headers to request
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, Folder_ID, User_ID")
		//rw.Header().Set("TESTINGGGGGG", "I'm here!!")
	}
	if req.Method == "OPTIONS" { // Stop here if its Preflighted OPTIONS request
		return
	}
	db.LogRequest(req.Method, req.URL.Path) // TEST: log url and request type
	CheckCookie(rw, req)                    // check for cookie

	c.R.ServeHTTP(rw, req) // Continue request
}
