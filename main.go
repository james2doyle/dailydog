// dailydog
//
// This is the main package which starts up and runs our REST API service.

package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

// main launches our web server which runs indefinitely.
func main() {

	slackWebhook := os.Getenv("SLACK_WEBHOOK")
	if slackWebhook == "" {
		log.Println("Error: you need to assign a `SLACK_WEBHOOK` environment variable.")
		os.Exit(1)
	}

	// Setup all routes.  We only service API requests, so this is basic.
	router := httprouter.New()
	router.GET("/", HandleIndex)

	// Setup 404 / 405 handlers.
	router.NotFound = http.HandlerFunc(NotFound)
	router.MethodNotAllowed = http.HandlerFunc(MethodNotAllowed)
	// router.PanicHandler = PanicHandler

	// Setup middlewares
	handler := cors.Default().Handler(router)

	// Start the server.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Starting HTTP server on port:", port)
	log.Fatal(http.ListenAndServe(":" + port, handler))

}
