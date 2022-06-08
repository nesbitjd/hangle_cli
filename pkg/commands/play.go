package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hangle_cli/pkg/hangman"
	"io/ioutil"
	"log"
	"net/http"
)

type SourceHangman struct {
	Word     string
	Failures int
	Guesses  []string
}

func play() {
	// Contact server for word
	id := 1
	url := fmt.Sprintf("http://127.0.0.1:8080/api/v1/hangman/%+v", id)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}

	game := &SourceHangman{}

	err = json.Unmarshal(body, &game)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(game)

	h := hangman.Init(game.Word)

	// Run game to completion
	if h.Play() {
		fmt.Println("You won!")
		game.uploadResults(h, url)
		return
	}

	fmt.Println("You suck.")
	game.uploadResults(h, url)
	// Upload results to server

}

func (game *SourceHangman) uploadResults(h *hangman.Hangman, url string) error {
	game.Failures = h.Failures
	game.Guesses = h.Guesses
	fmt.Println(*game)

	putBody, err := json.Marshal(game)
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(putBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	fmt.Println(resp.StatusCode)
	return nil
}
