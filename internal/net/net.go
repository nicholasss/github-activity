// Package net is for performing the requests and providing the responses
package net

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/nicholasss/github-activity/internal/data"
)

// === Constants === //

const GitHubAPIBaseURL string = "https://api.github.com/"

// === Variables === //
//
// empty

// === Intermodule Functions === //

// buildGetUserEventsURL can eventually add
func buildGetUserEventsURL(username string) (*url.URL, error) {
	URL, err := url.Parse(GitHubAPIBaseURL)
	if err != nil {
		return nil, err
	}

	// should make the following username
	// "https://api.github.com/users/nicholasss/events"
	URL = URL.JoinPath("users", username, "events")
	// fmt.Printf("DEBUG: url is %q\n", URL.String())
	return URL, nil
}

// addGitHubHeaders will directly add the necessary headers to the request
func addGitHubHeaders(req *http.Request) {
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Add("User-Agent", "nicholasss-github-activity-cli")
}

// checkResponseStatus will only return an error if there is a non-OK status
func checkResponseStatus(status int) error {
	if status >= 200 && status <= 299 {
		return nil
	}
	// status is not-OK, but what is it specifically?
	if status >= 100 && status <= 199 {
		return fmt.Errorf("information still processing.. no response? %s", http.StatusText(status))
	} else if status >= 300 && status <= 399 {
		return fmt.Errorf("performing redirect.. %s", http.StatusText(status))
	} else if status >= 400 && status <= 499 {
		return fmt.Errorf("client error.. %s", http.StatusText(status))
	} else if status >= 500 && status <= 599 {
		return fmt.Errorf("server error.. %s", http.StatusText(status))
	}

	return fmt.Errorf("http error of %s", http.StatusText(status))
}

// === Exported Functions === //

// FetchUserEvents will perform the request, format the response, and print the list of activity
func FetchUserEvents(username string) ([]data.GithubEvent, error) {
	// constructing url
	reqURL, err := buildGetUserEventsURL(username)
	if err != nil {
		return nil, err
	}

	// creating request
	//
	// body means the body of the req, so it would be used for POST or PUT
	//    to upload data to the server
	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	addGitHubHeaders(req)

	// performing request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// check the error
	err = checkResponseStatus(res.StatusCode)
	if err != nil {
		return nil, err
	}

	// TODO: eventually add checking for rate limit headers
	// e.g. `x-ratelimit-limit`, `x-ratelimit-reset`

	events, err := data.Decode(&res.Body)
	if err != nil {
		return nil, err
	}

	// return events to be formatted elsewhere
	return events, nil
}
