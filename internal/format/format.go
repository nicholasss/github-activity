// Package format will be responsible for formatting github events for the terminal.
package format

import (
	"fmt"

	"github.com/nicholasss/github-activity/internal/data"
)

// === Constants === //
//
// empty

// === Variables === //

// singleEventFormatStrings maps the event type to a premade single event format string
var singleEventFormatStrings map[string]string = map[string]string{
	"CommitCommentEvent":            " - NO FORMAT STRING YET\n",
	"CreateEvent":                   " - Created %s in %s\n",
	"DeleteEvent":                   " - NO FORMAT STRING YET\n",
	"ForkEvent":                     " - NO FORMAT STRING YET\n",
	"GollumEvent":                   " - NO FORMAT STRING YET\n",
	"IssueCommentEvent":             " - Commented in %s\n",
	"IssuesEvent":                   " - NO FORMAT STRING YET\n",
	"MemberEvent":                   " - NO FORMAT STRING YET\n",
	"PublicEvent":                   " - NO FORMAT STRING YET\n",
	"PullRequestEvent":              " - NO FORMAT STRING YET\n",
	"PullRequestReviewEvent":        " - NO FORMAT STRING YET\n",
	"PullRequestReviewCommentEvent": " - NO FORMAT STRING YET\n",
	"PullRequestReviewThreadEvent":  " - NO FORMAT STRING YET\n",
	"PushEvent":                     " - Pushed commit to %s\n",
	"ReleaseEvent":                  " - NO FORMAT STRING YET\n",
	"SponsorshipEvent":              " - NO FORMAT STRING YET\n",
	"WatchEvent":                    " - Starred %s\n",
}

// multiEventFormatStrings maps the event type to a premade multiple event format string
// the number will always be first, then the repo name
var multiEventFormatStrings map[string]string = map[string]string{
	"CommitCommentEvent":            " - NO FORMAT STRING YET\n",
	"CreateEvent":                   " - Created %s in %s\n",
	"DeleteEvent":                   " - NO FORMAT STRING YET\n",
	"ForkEvent":                     " - NO FORMAT STRING YET\n",
	"GollumEvent":                   " - NO FORMAT STRING YET\n",
	"IssueCommentEvent":             " - Commented %d times in %s\n",
	"IssuesEvent":                   " - NO FORMAT STRING YET\n",
	"MemberEvent":                   " - NO FORMAT STRING YET\n",
	"PublicEvent":                   " - NO FORMAT STRING YET\n",
	"PullRequestEvent":              " - NO FORMAT STRING YET\n",
	"PullRequestReviewEvent":        " - NO FORMAT STRING YET\n",
	"PullRequestReviewCommentEvent": " - NO FORMAT STRING YET\n",
	"PullRequestReviewThreadEvent":  " - NO FORMAT STRING YET\n",
	"PushEvent":                     " - Pushed %d commits to %s\n",
	"ReleaseEvent":                  " - NO FORMAT STRING YET\n",
	"SponsorshipEvent":              " - NO FORMAT STRING YET\n",
	"WatchEvent":                    " - Starred %s\n",
}

// === Intermodule Functions === //
//
// empty

// === Exported Functions === //

// PrintEvents prints out the events as a formatted list
func PrintEvents(events []data.GithubEvent) error {
	// switch statment for each event to print the relevant information
	//
	// how to combine multiple types in a row into a single one with a number?
	//   - Maybe if the repo & event.type is the same for the next one, we can combine and use new format?
	//
	// actually print out the strings

	// first loop through the event types and append them to a new list
	// eventTypes := make([]string, 0)
	// for _, event := range events {
	// 	eventTypes = append(eventTypes, event.Type)
	// }

	// we want to count how many of event there are in a row of the same type and repo
	// so that we can aggregate them into one line
	// e.g. "3 pushes were made to user/repo"
	//
	// eventSummary will have the latest event string
	// eventNumbers will have how many of the event occured
	// The length of the array should match

	if len(events) == 0 {
		return fmt.Errorf("unable to format empty list of events")
	}

	eventFormatStrs := make([]string, 0)
	eventRepos := make([]string, 0)
	eventCounts := make([]int, 0)

	previousEventType := ""
	previousEventRepo := ""
	for _, event := range events {
		eventType := event.Type
		eventRepo := event.Repo.Name

		// fmt.Printf("DEBUG: looking at %q in %q\n", eventType, eventRepo)
		// fmt.Printf("DEBUG: previous was %q in %q\n", previousEventType, previousEventRepo)

		// is the current the same as the previous? just modify the previous
		if previousEventType == "" || previousEventRepo == "" {
			eventFormatStrs = append(eventFormatStrs, singleEventFormatStrings[eventType])
			eventRepos = append(eventRepos, eventRepo)
			eventCounts = append(eventCounts, 1)

		} else if eventType == previousEventType && eventRepo == previousEventRepo {
			// this event type and repo is the same as the previous
			//
			// first check if the count is 1 or 2.
			//    if its 1 then replace the summary string **and** increment the last element
			//    if its 2 then only increment the last element
			//
			// either way, the name of the repo does not change

			currentEventCount := eventCounts[len(eventCounts)-1]
			if currentEventCount <= 1 {
				// replace the last string with the new string and increment
				eventFormatStrs[len(eventFormatStrs)-1] = multiEventFormatStrings[eventType]
			}
			// number is always incremented since its the same as the last event & repo
			eventCounts[len(eventCounts)-1] += 1

		} else {
			// event & repo were not the same as the last event & repo
			//
			// new format string for eventSummary
			// new repo name for eventRepos
			// new count for eventNumbers

			eventFormatStrs = append(eventFormatStrs, singleEventFormatStrings[eventType])
			eventRepos = append(eventRepos, eventRepo)
			eventCounts = append(eventCounts, 1)
		}

		// setup for next loop
		previousEventType = eventType
		previousEventRepo = eventRepo
	}

	// debug checking
	if len(eventCounts) == len(eventRepos) && len(eventRepos) == len(eventFormatStrs) && len(eventCounts) != 0 {
		// fmt.Printf("DEBUG: all event lists are the same length\n")
	} else {
		fmt.Printf("DEBUG: count problem!\n - summary: %d, repos: %d, numbers: %d\n", len(eventFormatStrs), len(eventRepos), len(eventCounts))
	}

	// printing out the format strings:
	for i, eventFormatStr := range eventFormatStrs {
		eventRepo := eventRepos[i]
		eventCount := eventCounts[i]

		if eventCount == 1 {
			fmt.Printf(eventFormatStr, eventRepo)
		} else {
			fmt.Printf(eventFormatStr, eventCount, eventRepo)
		}
	}

	return nil
}
