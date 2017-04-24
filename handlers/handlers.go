// dailydog
//
// The handlers for the routes

package handlers

import (
	"fmt"
	"github.com/james2doyle/dailydog/webhook"
	"github.com/julienschmidt/httprouter"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Setup the environment
	dogJson := os.Getenv("DOG_JSON")
	if dogJson == "" {
		log.Println("Info: using the default GIPHY API to fetch a random dog.")
		dogJson = "https://api.giphy.com/v1/gifs/random?api_key=dc6zaTOxFJmzC&tag=dog"
	}

	slackWebhook := os.Getenv("SLACK_WEBHOOK")
	if slackWebhook == "" {
		log.Println("Error: you need to assign a `SLACK_WEBHOOK` environment variable.")
	}

	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get(dogJson)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	value := gjson.GetBytes(body, "data.image_url")

	status := webhook.Post(true, slackWebhook, value.String())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(status))
}

// MethodNotAllowed renders a method not allowed response for invalid request
// types.
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	resp := webhook.Panic("Method Not Allowed")
	fmt.Fprintf(w, string(resp))
}

// NotFound renders a not found response for invalid API endpoints.
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	resp := webhook.Panic("Not Found")
	fmt.Fprintf(w, string(resp))
}

func PanicHandler(w http.ResponseWriter, r *http.Request, rcv interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	log.Println("Panic:", rcv)
	resp := webhook.Panic(rcv)
	fmt.Fprintf(w, string(resp))
}
