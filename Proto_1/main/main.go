package main

import (
	"log"
	"github.com/go-openapi/loads"
)

func main() {

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
	  log.Fatalln(err)
	}
	api := operations.NewHelloAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()
	server.Port = 8080
	// Start listening using having the handlers and port
	// already set up.
	if err := server.Serve(); err != nil {
	  log.Fatalln(err)
	}
  

}
