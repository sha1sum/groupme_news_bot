/*
Package bot handles posting a message to a GroupMe bot and handling the parsing of incoming messages from callbacks.

To use the bot functionality, you will need to first set BotID to the ID of the bot you wish to use.
*/
package bot

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// BotID is the ID of the GroupMe bot as found on GroupMe's developer site.
var BotID string

type (
	// IncomingMessage is used to indicate the message properties from the POST sent from a GroupMe bot callback.
	IncomingMessage struct {
		// Attachments is a list of attachments added on to the GroupMe post
		Attachments []Attachment `json:"attachments"`
		// AvatarURL is the address of the avatar used by the poster
		AvatarURL string `json:"avatar_url"`
		// CreatedAt is the unix timestamp when the message was posted
		CreatedAt uint32 `json:"created_at"`
		// GroupID is the ID of the GroupMe group that the message was posted to
		GroupID string `json:"group_id"`
		// ID is the message ID for the post
		ID string `json:"id"`
		// Name is the nickname of the person/bot which posted the message
		Name   string `json:"name"`
		// SenderID is a GroupMe internal ID for the sender of the post
		SenderID string `json:"sender_id"`
		// SenderType is the type of member sending the post (bot or user)
		SenderType string `json:"sender_type"`
		// SenderGUID is a global unique identifier for the poster
		SenderGUID string `json:"sender_guid"`
		// System indicates whether GroupMe's server sent the message or whether it was initiated by a poster
		System bool `json:"system"`
		// Text is the message text that was posted
		Text   string `json:"text"`
		// UserID is the GroupMe User ID of the person/bot which posted the message
		UserID string `json:"user_id"`
	}

	// Attachment is a type of embedded attachment for a GroupMe post
	Attachment struct {
		// Loci is a list of the optional starts and ends of mentions in the string where "@Test User" would be [0, 9]
		Loci [][2]int `json:"loci"`
		// Type is the type of attachment ("image", "video", "location", "event", or "mentions")
		Type string `json:"type"`
		// UserIDs is an optional array of GroupMe User IDs mentioned in the message text
		UserIDs []int `json:"user_ids"`
		// PreviewURL is a URL for a snapshot of an attached video
		PreviewURL string `json:"preview_url"`
		// URL is the URL for an image or video attachment
		URL string `json:"url"`
		// Lat is the latitude for an attached location
		Lat string `json:"lat"`
		// Lng is the longitude for an attached location
		Lng string `json:"lng"`
		// Name is the name of an attached location
		Name string `json:"name"`
		// EventID is the GroupMe internal identifier for a created calendar event
		EventID string `json:"event_id"`
		// View is the type of preview to attach for a calendar event
		View string `json:"view"`
	}

	// OutgoingMessage houses a string message along with any error that may have resulted from running a Handler.
	OutgoingMessage struct {
		Message string
		Err     error
	}

	// Handler will be used to perform actions given an IncomingMessage and output an OutgoingMessage result to a channel.
	Handler interface {
		Handle(term string, c chan []*OutgoingMessage, message IncomingMessage)
	}
)

// PostMessage posts a string to a GroupMe bot as long as the BotID is present.
func PostMessage(message string) (*http.Response, error) {
	if len(BotID) < 1 {
		return nil, errors.New("BotID cannot be blank.")
	}
	m := map[string]string{"bot_id": BotID, "text": message}
	j, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return http.Post("https://api.groupme.com/v3/bots/post", "application/json", strings.NewReader(string(j)))
}
