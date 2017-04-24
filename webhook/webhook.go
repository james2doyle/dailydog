// dailydog
//
// The general function to post to Slack

package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/james2doyle/dailydog/models"
	"io/ioutil"
	"net/http"
)

func Post(status bool, hook, messageAddon string) []byte {
	var message string
	if status {
		message = fmt.Sprintf("*Woof!* Here is your daily dog!\n<%s|View This GIF>", messageAddon)
	} else {
		message = fmt.Sprintf("*Oops!* There was an error fetching your dailydog.\n%s", messageAddon)
	}

	data := models.Webhook{
		Username: "Daily Dog",
		IconUrl:  "https://i.imgur.com/0Uzt9VB.png",
		Text:     message,
	}

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(data)

	resp, err := http.Post(hook, "application/json; charset=utf-8", buffer)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	result := models.SlackResponse{string(body)}

	jsonResponse, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return jsonResponse
}

func Panic(message interface{}) []byte {
	result := models.PanicResponse{fmt.Sprintf("%s", message)}

	jsonResponse, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return jsonResponse
}
