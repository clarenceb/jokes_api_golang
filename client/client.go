package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to the Jokes API client!")

	resp, err := http.Get("http://localhost:8080/joke")

	if err != nil {
		fmt.Println(err)
		return
	}

	var body struct {
		Joke string `json:"joke"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(body.Joke)
}
