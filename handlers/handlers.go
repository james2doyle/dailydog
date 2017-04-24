// dailydog
//
// The handlers for the routes

package handlers

import (
  "fmt"
  "log"
  "github.com/james2doyle/dailydog/endpoint"
  "github.com/julienschmidt/httprouter"
  "github.com/tidwall/gjson"
  "io/ioutil"
  "net/http"
  "time"
  "os"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")

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
    MaxIdleConns:       10,
    IdleConnTimeout:    30 * time.Second,
  }

  client := &http.Client{Transport: tr}
  resp, err := client.Get(dogJson)
  if err != nil {
    handleError(w, slackWebhook, err)
  }

  defer resp.Body.Close()

  if err != nil {
    handleError(w, slackWebhook, err)
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    handleError(w, slackWebhook, err)
  }

  value := gjson.GetBytes(body, "data.image_url")

  status := endpoint.Post(true, slackWebhook, value.String())

  fmt.Fprintf(w, status)
}

// MethodNotAllowed renders a method not allowed response for invalid request
// types.
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusMethodNotAllowed)
}

// NotFound renders a not found response for invalid API endpoints.
func NotFound(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusNotFound)
}

func handleError(w http.ResponseWriter, slackWebhook string, err error) {
  log.Println(err)
  panic(err)
  w.WriteHeader(http.StatusInternalServerError)
  status := endpoint.Post(false, slackWebhook, err.Error())
  fmt.Fprintf(w, status)
}
