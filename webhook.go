// dailydog
//
// The general function to post to Slack

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// WebhookData -- the struct for the JSON that is sent to the Slack endpoint
type WebhookData struct {
	Username string `json:"username"` // "Daily Dog"
	IconURL  string `json:"icon_url"` // "https://i.imgur.com/0Uzt9VB.png"
	Text     string `json:"text"`     // "<https://i.imgur.com/0Uzt9VB.png|View Photo>\nThis is a line of text in a channel."
}

// SlackResponse -- structure of the JSON response from the Slack Webhook API
type SlackResponse struct {
	Status string `json:"status"`
}

// WebhookPost -- the function that handles the call to the Slack Webhook
func WebhookPost(status bool, hook, messageAddon string) SlackResponse {
	var message string
	if status {
		message = fmt.Sprintf("*Woof!* Here is your daily dog!\n<%s|View This GIF>", messageAddon)
	} else {
		message = fmt.Sprintf("*Oops!* There was an error fetching your dailydog.\n%s", messageAddon)
	}

	data := WebhookData{
		Username: "Daily Dog",
		IconURL:  "https://i.imgur.com/0Uzt9VB.png",
		Text:     message,
	}

	output, err := json.Marshal(data)
	if err != nil {
		return SlackResponse{Status: err.Error()}
	}

	buffer := bytes.NewReader(output)

	resp, err := http.Post(hook, "application/json; charset=utf-8", buffer)
	if err != nil {
		return SlackResponse{Status: err.Error()}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SlackResponse{Status: err.Error()}
	}

	result := SlackResponse{Status: string(body)}

	return result
}
