package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Jokes API client
type jokeApi struct {
	BaseUri     string
	NumRequests int
}

// Fetch jokes synchronously: one at a time
func (api jokeApi) SyncJokes() []string {
	results := make([]string, api.NumRequests)

	for i := 0; i < api.NumRequests; i++ {
		joke, err := api.joke()

		if err != nil {
			log.Println(err)
		} else {
			results[i] = fmt.Sprintf("%d: %s", i+1, joke)
		}
	}

	return results
}

// Fetch jokes concurrently: many at a time
func (api jokeApi) ConcurrentJokes() []string {
	jokes := make(chan string, api.NumRequests)
	errs := make(chan error, api.NumRequests)

	for i := 0; i < api.NumRequests; i++ {
		go func(i int) {
			joke, err := api.joke()
			if err != nil {
				errs <- err
				return
			}
			jokes <- fmt.Sprintf("%d: %s", i+1, joke)
		}(i)
	}

	results := make([]string, api.NumRequests)

	for i := 0; i < api.NumRequests; i++ {
		select {
		case joke := <-jokes:
			results[i] = joke
		case err := <-errs:
			log.Printf(err.Error())
		}
	}

	return results
}

// Fetch a single joke from an external HTTP web service
func (api jokeApi) joke() (string, error) {
	resp, err := http.Get(api.BaseUri + "/joke")

	if err != nil {
		log.Println(err)
		return "", err
	}

	var body struct {
		Joke string `json:"joke"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		log.Println(err)
		return "", err
	}

	return body.Joke, nil
}
