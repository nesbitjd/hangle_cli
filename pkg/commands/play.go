package commands

import (
	"fmt"
	"hangle_cli/pkg/hangman"
	"log"
	"strings"

	"github.com/nesbitjd/hangle_server/types"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

var (
	hangleServerUrl = "http://127.0.0.1:8080"
)

func play() {
	// Contact server for word

	word, err := types.GetLastWord(hangleServerUrl)
	if err != nil {
		log.Fatalln(err)
		return
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

	u, err := checkUser(username, hangleServerUrl)
	if err != nil {
		log.Fatalln(err)
		return
	}

	if u.Username == "" {
		u.Username = username

		u.PostUser(hangleServerUrl)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}

	record := types.NewRecord(word, *u, h.Failures, strings.Join(h.Guesses, ","))

	// Upload results to server
	record.PostResults(hangleServerUrl)
}

// checkUser takes a lookupUser to look up, and the base_url for the api.
// If found, returns a pointer to the user. Otherwise returns empty user
func checkUser(lookupUser string, base_url string) (*types.User, error) {
	users, err := types.GetAllUsers(hangleServerUrl)
	if err != nil {
		return nil, xerrors.Errorf("unable to get users: %w", err)
	}

	for _, u := range users {
		if u.Username == lookupUser {
			logrus.Infof("user %q found", lookupUser)
			return &u, nil
		}
	}

	logrus.Infof("user %q not found", lookupUser)
	u := &types.User{}
	return u, nil
}
