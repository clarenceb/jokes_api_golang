package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Jokes API!"))
	})

	http.HandleFunc("/joke", joke)

	http.ListenAndServe(":8080", nil)
}

// Handles /joke route requests.
func joke(w http.ResponseWriter, r *http.Request) {
	joke, err := randomJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"joke": joke,
	})
}

// Fetches a single random joke from ICNDb.com API.
func randomJoke() (string, error) {
	log.Printf("Fetching a random joke...")
	resp, err := http.Get("http://api.icndb.com/jokes/random")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var joke struct {
		Value struct {
			Joke string `json:"joke"`
		} `json:"value"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		return "", err
	}

	log.Printf("Joke is: %s", joke.Value.Joke)
	return joke.Value.Joke, nil
}
