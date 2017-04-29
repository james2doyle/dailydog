// dailydog
//
// This is the main package which starts up and runs our REST API service.

package main

import (
	"log"
	"net/http"
	"os"
)

// main launches our web server which runs indefinitely.
func main() {

	slackWebhook := os.Getenv("SLACK_WEBHOOK")
	if slackWebhook == "" {
		log.Fatal("\033[0;31mError: you need to assign a `SLACK_WEBHOOK` environment variable.\033[0m")
	}

	// Setup all routes.  We only service API requests, so this is basic.
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleIndex)

	// Start the server.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Starting HTTP server on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))

}
