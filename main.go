// dailydog
//
// This is the main package which starts up and runs our REST API service.

package main

import (
	"github.com/james2doyle/dailydog/handlers"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

// main launches our web server which runs indefinitely.
func main() {

	// Setup all routes.  We only service API requests, so this is basic.
	router := httprouter.New()
	router.GET("/", handlers.Index)

	// Setup 404 / 405 handlers.
	router.NotFound = http.HandlerFunc(handlers.NotFound)
	router.MethodNotAllowed = http.HandlerFunc(handlers.MethodNotAllowed)
	router.PanicHandler = handlers.PanicHandler

	// Setup middlewares
	handler := cors.Default().Handler(router)

	// Start the server.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Starting HTTP server on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}
