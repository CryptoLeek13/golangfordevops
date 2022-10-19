package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func main() {
	args := os.Args
	// Error handling incase of less then two argument
	// Responses provided
	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url>")
		os.Exit(1)
	}
	// Error handling to parse the url arguments to ensure they are
	// of correct format i.e url (string)
	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("Url is in invalid format: %s\n", err)
		os.Exit(1)
	}
	// HTTP.GET request to url to provide its status and body resonse
	response, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}
	// defer the closing of the response body until the function returns
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// update response.body byte stream into string format
	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP Code %d): %s\n", response.StatusCode, string(body))
		os.Exit(1)
	}
	var words Words

	err = json.Unmarshal(body, &words)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSON Parsed\nPage: %s\nWords: %v\n", words.Page, words.Words)
}
