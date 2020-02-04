package main

import (
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

	db.SetupDB()              // setup db tables and add sample data
	router := mux.NewRouter() // create router
	ServeRoutes(router)       // handles API Route requests

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
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, YourOwnHeader")
		//rw.Header().Set("TESTINGGGGGG", "I'm here!!")
	}

	if req.Method == "OPTIONS" { // Stop here if its Preflighted OPTIONS request
		return
	}

	CheckCookie(rw, req) // check for cookie

	db.LogRequest(req.Method, req.URL.Path) // TEST: log url and request type

	c.R.ServeHTTP(rw, req) // Continue request
}
