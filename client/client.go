package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to the Jokes API client!")

	var n = flag.Int("n", 1, "Number of jokes to get")
	var c = flag.Bool("c", false, "Fetch jokes concurrently")
	flag.Parse()

	if *c {
		concurrentJokes(*n)
	} else {
		syncJokes(*n)
	}
}

func syncJokes(n int) {
	for i := 0; i < n; i++ {
		joke, err := joke()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%d: %s\n", i+1, joke)
	}
}

func concurrentJokes(n int) {
	jokes := make(chan string, n)
	errs := make(chan error, n)

	for i := 0; i < n; i++ {
		go func(i int) {
			joke, err := joke()
			if err != nil {
				errs <- err
				return
			}
			jokes <- fmt.Sprintf("%d: %s\n", i+1, joke)
		}(i)
	}

	for i := 0; i < n; i++ {
		select {
		case joke := <-jokes:
			fmt.Println(joke)
		case err := <-errs:
			fmt.Printf(err.Error())
		}
	}
}

func joke() (string, error) {
	resp, err := http.Get("http://localhost:8080/joke")

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var body struct {
		Joke string `json:"joke"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		fmt.Println(err)
		return "", err
	}

	return body.Joke, nil
}
