package commands

import (
	"fmt"
	"github.com/nesbitjd/hangle_cli/pkg/hangman"
	"net/http"
	"strings"

	"github.com/nesbitjd/hangle_server/pkg/hangle"
)

var (
	hangleServerUrl = "http://127.0.0.1:8080"
)

func play() error {
	// Contact server for word
	client := hangle.NewClient(hangle.NewConfig(hangleServerUrl), http.DefaultClient)

	word, err := client.GetLastWord()
	if err != nil {
		return fmt.Errorf("unable to get last word, %w", err)
	}

	h := hangman.Init(word.Word)

	// Run game to completion
	if h.Play() {
		fmt.Println("You won!")
	} else {
		fmt.Println("You suck.")
	}

	// Retrieve user and validate input

	var username string
	for !func(u string) bool {
		valid := true
		fmt.Println("Enter name to upload results:")
		fmt.Scanln(&u)
		username = u

		if u == "" {
			fmt.Println("Invalid input: empty string")
			valid = false
		}

		return valid
	}(username) {
	}

	u, err := checkUser(username, client)
	if err != nil {
		return fmt.Errorf("unable to validate user, %w", err)
	}

	if u.Username == "" {
		u.Username = username

		client.PostUser(*u)
		if err != nil {
			return fmt.Errorf("unable to post user, %w", err)
		}
	}

	record := hangle.NewRecord(word, *u, h.Failures, strings.Join(h.Guesses, ","))

	// Upload results to server
	client.PostRecord(record)

	return nil
}

// checkUser takes a lookupUser to look up, and the base_url for the api.
// If found, returns a pointer to the user. Otherwise returns empty user
func checkUser(lookupUser string, client hangle.Client) (*hangle.User, error) {
	users, err := client.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("unable to get users: %w", err)
	}

	for _, u := range users {
		if u.Username == lookupUser {
			return &u, nil
		}
	}

	u := &hangle.User{}
	return u, nil
}
