package hangman

import (
	"errors"
	"strings"
)

type Hangman struct {
	wordToGuess string
	guessing    Guessing
}

type Guessing struct {
	WordGuessedSoFar string
	IsGuessed        bool
	Fails            int
	LettersAttempted []rune
}

func (h *Hangman) SetWord(w string) {
	h.wordToGuess = strings.ToLower(w)
	h.guessing.WordGuessedSoFar = obfuscate(w)
}

func (h Hangman) GetObfuscatedWord() string {
	return h.guessing.WordGuessedSoFar
}

func (h Hangman) GetGuessing() Guessing {
	return h.guessing
}

func (h *Hangman) Guess(guess string) (Guessing, error) {
	if len(guess) > 1 {
		return Guessing{}, errors.New("you can guessonly one letter")
	}

	var letterFound bool
	wordGuessedSoFar := []rune(h.guessing.WordGuessedSoFar)
	h.guessing.IsGuessed = false
	h.guessing.LettersAttempted = appendAttemptedLetterOnce(h.guessing.LettersAttempted, guess)

	for i, letter := range h.wordToGuess {
		if string(letter) == strings.ToLower(guess) {
			wordGuessedSoFar[i] = letter
			letterFound = true
		}
	}

	h.guessing.WordGuessedSoFar = string(wordGuessedSoFar)

	if !letterFound {
		h.guessing.Fails++
	}

	if h.guessing.WordGuessedSoFar == h.wordToGuess {
		h.guessing.IsGuessed = true
	}

	return h.guessing, nil
}

func obfuscate(w string) string {
	var encrypted string
	for _, word := range w {
		if word == ' ' {
			encrypted += " "
			continue
		}
		encrypted += "_"
	}
	return encrypted
}

func appendAttemptedLetterOnce(originalSlice []rune, guess string) []rune {
	for _, r := range originalSlice {
		if r == rune(guess[0]) {
			return originalSlice
		}
	}
	return append(originalSlice, rune(guess[0]))
}
