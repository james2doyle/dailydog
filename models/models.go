// dailydog/models
//
// This package contains all models used in the dailydog service.

package models

// Webhook is a struct we use to represent JSON API responses.
type Webhook struct {
  Username string `json:"username"` // "Daily Dog"
  IconUrl string `json:"icon_url"` // "https://i.imgur.com/0Uzt9VB.png"
  Text string `json:"text"` // "<https://i.imgur.com/0Uzt9VB.png|View Photo>\nThis is a line of text in a channel."
}
