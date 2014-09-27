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
	flag.Parse()

	for i := 0; i < *n; i++ {
		joke, err := joke()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%d: %s\n", i+1, joke)
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
