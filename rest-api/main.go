package main

import (
	"FYP_Proto_Backend/api"
	m "FYP_Proto_Backend/model"
	"fmt"
	"net/http"

	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

func main() {

	// mock data for testing
	m.Folders = append(m.Folders, m.Folder{ID: "1", Name: "Security Notes", Notes: "Testing 123 ABC"})
	m.Folders = append(m.Folders, m.Folder{ID: "2", Name: "Computer Architecture", Notes: "Testing 124"})

	router := mux.NewRouter()

	/*
	 * FOLDERS
	 */
	router.HandleFunc("/folders", api.GetFolders).Methods("GET", "OPTIONS")
	router.HandleFunc("/folder", api.CreateFolder).Methods("POST", "OPTIONS")
	router.HandleFunc("/folders/{id}", api.GetFolder).Methods("GET", "OPTIONS")
	router.HandleFunc("/folders/{id}", api.UpdateFolderName).Methods("PUT", "OPTIONS")
	router.HandleFunc("/folders/{id}", api.DeleteFolder).Methods("DELETE", "OPTIONS")

	// Run server
	port := ":8011"
	fmt.Println("Started Server on port", port)
	//http.ListenAndServe(port, router)
	http.ListenAndServeTLS(":8011", "localhost.crt", "localhost.key", &CORSRouterDecorator{router})

	//certmagic.HTTPS([]string{"localhost:8010"}, router)

}

// ServeHTTP wraps the HTTP server enabling CORS headers.
// For more info about CORS, visit https://www.w3.org/TR/cors/
func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, YourOwnHeader")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}
