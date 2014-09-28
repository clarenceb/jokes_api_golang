package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Jokes API client!")

	var n = flag.Int("n", 1, "Number of jokes to get")
	var c = flag.Bool("c", false, "Fetch jokes concurrently")
	flag.Parse()

	api := jokeApi{BaseUri: "http://localhost:8080", NumRequests: *n}

	var jokes []string

	if *c {
		jokes = api.ConcurrentJokes()
	} else {
		jokes = api.SyncJokes()
	}

	for _, joke := range jokes {
		fmt.Println(joke)
	}
}
