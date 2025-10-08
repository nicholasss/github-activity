// Package format will be responsible for formatting github events for the terminal.
package format

import (
	"fmt"

	"github.com/nicholasss/github-activity/internal/data"
)

// === Types === //

type FormatEvent struct {
	EventType           string
	Count               int
	Repo                string
	CreateEventModifier string
}

// === Constants === //
//
// empty

// === Variables === //

// singleEventFormatStrings maps the event type to a premade single event format string
var singleEventFormatStrings map[string]string = map[string]string{
	"CommitCommentEvent":            " - Commit comment in %s\n",
	"CreateEvent":                   " - Created %s in %s\n",
	"DeleteEvent":                   " - Deleted branch/tag in %s\n",
	"ForkEvent":                     " - Forked repo to %s\n",
	"GollumEvent":                   " - Created or updated wiki page in %s\n",
	"IssueCommentEvent":             " - Commented in %s\n",
	"IssuesEvent":                   " - Issues action in %s\n",
	"MemberEvent":                   " - Collaborator action in %s\n",
	"PublicEvent":                   " - Changed %s to public\n",
	"PullRequestEvent":              " - Pull request activity in %s\n",
	"PullRequestReviewEvent":        " - Pull request review activity in %s\n",
	"PullRequestReviewCommentEvent": " - Pull request review comment activity in %s\n",
	"PullRequestReviewThreadEvent":  " - Pull request review thread activity in %s\n",
	"PushEvent":                     " - Pushed commit to %s\n",
	"ReleaseEvent":                  " - Release activity in %s\n",
	"SponsorshipEvent":              " - Sponsoship activity in %s\n",
	"WatchEvent":                    " - Starred %s\n",
}

// multiEventFormatStrings maps the event type to a premade multiple event format string
// the number will always be first, then the repo name
var multiEventFormatStrings map[string]string = map[string]string{
	"CommitCommentEvent":            " - %d commit comments in %s\n",
	"CreateEvent":                   " - Created %d %s's in %s\n",
	"DeleteEvent":                   " - Deleted %d branches/tags in %s\n",
	"ForkEvent":                     " - Forked %d repos to %s\n",
	"GollumEvent":                   " - Created or updated %d wiki pages in %s\n",
	"IssueCommentEvent":             " - Commented %d times in %s\n",
	"IssuesEvent":                   " - %d issue actions in %s\n",
	"MemberEvent":                   " - %d collaborator actions in %s\n",
	"PublicEvent":                   " - %d times changed %s into public\n",
	"PullRequestEvent":              " - %d pull request activities in %s\n",
	"PullRequestReviewEvent":        " - %d pull request review activities in %s\n",
	"PullRequestReviewCommentEvent": " - %d pull request review comment activities in %s\n",
	"PullRequestReviewThreadEvent":  " - %d pull request thread activities in %s\n",
	"PushEvent":                     " - Pushed %d commits to %s\n",
	"ReleaseEvent":                  " - %d release activities in %s\n",
	"SponsorshipEvent":              " - %d sponsorship activities in %s\n",
	"WatchEvent":                    " - Starred %s\n",
}

// === Intermodule Functions === //

func parseIntoFormatEvents(events []data.GithubEvent) []FormatEvent {
	previousEventType := ""
	previousRepo := ""
	previousCreateEventModifier := ""

	// list of new format events being created
	formatEvents := make([]FormatEvent, 0)

	for _, event := range events {
		// first bring over information into variables
		eventType := event.Type
		repo := event.Repo.Name
		createEventModifier := event.CreateEventType // should be "" when its not create event
		newFormatEvent := FormatEvent{}

		if previousEventType == eventType && previousRepo == repo && eventType != "CreateEvent" {
			// first check previous event before creating a new one, and current is not "CreateEvent"
			formatEvents[len(formatEvents)-1].Count += 1
		} else if eventType == "CreateEvent" && createEventModifier == previousCreateEventModifier {
			// then check for create events that are the same as the previous
			formatEvents[len(formatEvents)-1].Count += 1
		} else {
			// create new format event
			newFormatEvent.Count = 1
			newFormatEvent.EventType = eventType
			newFormatEvent.Repo = repo
			newFormatEvent.CreateEventModifier = createEventModifier // again should be "" when its not create event

			formatEvents = append(formatEvents, newFormatEvent)
		}

		// end of loop
		previousEventType = eventType
		previousRepo = repo
		previousCreateEventModifier = createEventModifier
	}

	return formatEvents
}

func printFormatEvents(events []FormatEvent) {
	for _, event := range events {
		if event.Count <= 1 {
			// single event
			formatStr := singleEventFormatStrings[event.EventType]

			if event.EventType == "CreateEvent" {
				fmt.Printf(formatStr, event.CreateEventModifier, event.Repo)
			} else {
				fmt.Printf(formatStr, event.Repo)
			}
		} else {
			// multiple count event
			formatStr := multiEventFormatStrings[event.EventType]

			if event.EventType == "CreateEvent" {
				fmt.Printf(formatStr, event.Count, event.CreateEventModifier, event.Repo)
			} else {
				fmt.Printf(formatStr, event.Count, event.Repo)
			}
		}
	}
}

// === Exported Functions === //

// PrintEvents prints out the events as a formatted list
func PrintEvents(events []data.GithubEvent) error {
	if len(events) == 0 {
		return fmt.Errorf("unable to format empty list of events")
	}
	// fmt.Printf("DEBUG: Original events:\n")
	// for _, event := range events {
	// 	if event.Type == "CreateEvent" {
	// 		fmt.Printf("DEBUG: - %v - %v - %v\n", event.Type, event.CreateEventType, event.Repo.Name)
	// 	}
	// 	fmt.Printf("DEBUG: - %v - %v\n", event.Type, event.Repo.Name)
	// }

	formatEvents := parseIntoFormatEvents(events)
	printFormatEvents(formatEvents)
	return nil
}
