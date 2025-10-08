// Package data provides data models for the rest of the program
package data

import (
	"encoding/json"
	"io"
	"time"
)

type GithubEvent struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Actor struct {
		ID           int    `json:"id"`
		Login        string `json:"login"`
		DisplayLogin string `json:"display_login"`
		GravatarID   string `json:"gravatar_id"`
		URL          string `json:"url"`
		AvatarURL    string `json:"avatar_url"`
	} `json:"actor"`
	Repo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
	Org       struct {
		ID         int    `json:"id"`
		Login      string `json:"login"`
		GravatarID string `json:"gravatar_id"`
		URL        string `json:"url"`
		AvatarURL  string `json:"avatar_url"`
	} `json:"org"`
	Payload         json.RawMessage `json:"payload"`
	CreateEventType string
}

// CreateEventPayload is for the specific 'createEvent'
type CreateEventPayload struct {
	RefType string `json:"ref_type"`
}

func Decode(responseBody *io.ReadCloser) ([]GithubEvent, error) {
	var events []GithubEvent
	err := json.NewDecoder(*responseBody).Decode(&events)
	if err != nil {
		return nil, err
	}

	for i, event := range events {
		switch event.Type {
		case "CreateEvent":
			var createEventPayload CreateEventPayload
			err := json.Unmarshal(event.Payload, &createEventPayload)
			if err != nil {
				return nil, err
			}
			// write decoded event direct to the original payload
			events[i].CreateEventType = createEventPayload.RefType

		default:
			continue
		}
	}

	return events, nil
}
