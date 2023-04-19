package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"modok64/hangman/hangman"
	"modok64/hangman/util"
	"net/http"
	"os"
	"time"
)

const (
	numberOfGuesses = 6
)

func main() {

	//wordToGuess := "congratulation at your new job by gabry"
	wordToGuess := downloadWord()
	hangman := hangman.New(wordToGuess)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		util.ClearScreen()
		fmt.Println(hangman.ObfuscatedWord())
		fmt.Println("Fails: " + fmt.Sprint(hangman.Guessing().Fails) + " out of " + fmt.Sprint(numberOfGuesses))
		fmt.Println("Letters tried so far: " + string(hangman.Guessing().LettersAttempted))

		fmt.Print("Guess the letter: ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) > 0 {
			guess, err := hangman.Guess(text)
			if err != nil {
				fmt.Println("You can guess one character at the time")
				time.Sleep(2 * time.Second)
				continue
			}
			if guess.IsGuessed {
				fmt.Println("You guessed it!")
				break
			}
			if guess.Fails == numberOfGuesses {
				fmt.Println("You lost!")
				fmt.Println("The word to guess was: " + wordToGuess)
				break
			}
		} else {
			fmt.Println("Exiting...")
			break
		}
	}
}

func downloadWord() string {
	var j []string
	res, err := http.Get("https://random-word-api.herokuapp.com/word")
	json.NewDecoder(res.Body).Decode(&j)

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	return j[0]
}
