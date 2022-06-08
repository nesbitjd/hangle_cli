package hangman

import (
	"fmt"
	"regexp"
	"strings"
)

type Hangman struct {
	Word      []Letter
	Guesses   []string
	Failures  int
	Successes int
}

type Letter struct {
	Value rune
	Mask  bool
}

const bodyParts = 6

// Play starts
func (h *Hangman) Play() bool {
	for h.Failures < bodyParts {
		fmt.Println(buildHangman(h.Failures))
		fmt.Println(h.toString())

		fmt.Println("Take a guess!")
		guess := getInput()

		err := h.handleInput(guess, h.Guesses)
		if err != nil {
			fmt.Printf("Improper input: %s\n", err)
			continue
		}

		if h.handleGuess(guess) {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Wrong!")
			continue
		}

		if checkWin(h) {
			fmt.Println(h.toString())
			return true
		}

	}
	fmt.Println(buildHangman(h.Failures))
	return false
}

// Init parses
func Init(w string) *Hangman {
	h := Hangman{}
	letter := Letter{}

	for _, l := range strings.ToLower(w) {
		letter = Letter{
			Value: l,
			Mask:  true,
		}
		h.Word = append(h.Word, letter)
	}

	h.Successes = 0
	h.Failures = 0
	h.Guesses = []string{}

	return &h
}

func (h *Hangman) toString() string {
	word := ""

	for _, l := range h.Word {
		if !l.Mask {
			word += string(l.Value)
		} else {
			word += "_ "
		}

	}

	return word
}

func getInput() string {
	var guess string
	fmt.Scanln(&guess)

	guess = strings.ToLower(guess)

	return guess
}

func (h *Hangman) handleInput(guess string, guesses []string) error {
	IsLetter := regexp.MustCompile(`^[a-zA-Z]$`).MatchString(guess)
	if !IsLetter {
		return fmt.Errorf("guess is not valid input")
	}

	for _, g := range guesses {
		if guess == g {
			return fmt.Errorf("guess has already been used")
		}
	}

	h.Guesses = append(h.Guesses, guess)

	return nil
}

func (h *Hangman) handleGuess(guess string) bool {
	found := false
	for i, r := range h.Word {
		//fmt.Printf("The rune is: %s, and the guess is: %s", strconv.QuoteRune(r.Value), guess)
		if string(r.Value) == guess {
			h.Successes += 1
			h.Word[i].Mask = false
			found = true
		}
	}
	if !found {
		h.Failures += 1
	}
	return found
}

func checkWin(h *Hangman) bool {
	return h.Successes >= len(h.Word)
}

func buildHangman(failures int) string {
	gallows := " __\n|  |"
	manParts := []string{"  O", " /", "|", "\\", " /", " \\"}

	for i := 0; i < 3; i++ {
		gallows = gallows + "\n|"

		for j := 0; j < failures; j++ {
			if i == 0 && j < 1 {
				gallows = gallows + manParts[j]
			}
			if i == 1 && j > 0 && j < 4 {
				gallows = gallows + manParts[j]
			}
			if i == 2 && j > 3 {
				gallows = gallows + manParts[j]
			}
		}
	}

	gallows = gallows + "\n|________"
	return gallows
}
