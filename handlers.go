// dailydog
//
// The handlers for the routes
// WriteJSON and JSONError from https://github.com/nesv/apiutil

package main

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// WriteJSON marshals the given value v to its JSON representation, and writes
// it to an http.ResponseWriter with the given HTTP status code. This function
// also makes sure to set the "Content-Type" header to "application/json".
func WriteJSON(w http.ResponseWriter, v interface{}, status int) {
	p, err := json.Marshal(&v)
	if err != nil {
		JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(p)
}

// JSONError is similar to http.Error() in where it allows you to write an
// error to an http.ResponseWriter with a given HTTP status, except for that
// it will wrap your error in a JSON object, and put the error message under
// the object's "error" key.
//
// Calling this function will result in a response body like so:
//
// 	{"error": "...your error message..."}
//
func JSONError(w http.ResponseWriter, errStr string, status int) {
	v := jsonResponse{Status: status, Message: errStr}
	WriteJSON(w, v, status)
}

// HandleIndex -- the main function for all routes
func HandleIndex(w http.ResponseWriter, req *http.Request) {
	// Setup the environment
	dogJSON := os.Getenv("DOG_JSON")
	if dogJSON == "" {
		log.Println("\033[0;36mInfo: using the default GIPHY API to fetch a random dog.\033[0m")
		dogJSON = "https://api.giphy.com/v1/gifs/random?api_key=dc6zaTOxFJmzC&tag=dog"
	}

	// we know this exists - because we got this far
	slackWebhook := os.Getenv("SLACK_WEBHOOK")

	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get(dogJSON)
	if err != nil {
		JSONError(w, err.Error(), http.StatusInternalServerError)
	}

	defer resp.Body.Close()

	if err != nil {
		JSONError(w, err.Error(), http.StatusInternalServerError)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		JSONError(w, err.Error(), http.StatusInternalServerError)
	}

	value := gjson.GetBytes(body, "data.image_url")

	status := WebhookPost(true, slackWebhook, value.String())

	WriteJSON(w, status, http.StatusOK)
}
