package hangman

import "testing"

func TestEncryptWithOneWord(t *testing.T) {
	encrypted := obfuscate("test")

	assertWordIs(encrypted, "____", t)

	encrypted = obfuscate("longwordtotest")

	assertWordIs(encrypted, "______________", t)
}

func TestEncryptWithTwoWords(t *testing.T) {
	encrypted := obfuscate("test test")
	assertWordIs(encrypted, "____ ____", t)
}

func TestEncryptWithThreeWords(t *testing.T) {
	encrypted := obfuscate("test test test")
	assertWordIs(encrypted, "____ ____ ____", t)
}

func TestWordIsCaseInsensitive(t *testing.T) {
	h := Hangman{}
	h.SetWord("HoLa")

	guess, _ := h.Guess("h")
	assertWordIs(guess.WordGuessedSoFar, "h___", t)

	secondGuess, _ := h.Guess("O")
	assertWordIs(secondGuess.WordGuessedSoFar, "ho__", t)
}

func TestGuessingMoreThanOneCharAtTimeShouldError(t *testing.T) {
	h := Hangman{}
	h.SetWord("HoLa")

	_, err := h.Guess("ho")

	if err == nil {
		t.Errorf("Expected error. More than one letter passed into the function")
	}
}

func TestGuessShouldReturnFailures(t *testing.T) {
	h := Hangman{}
	h.SetWord("hoLa")

	guess, _ := h.Guess("h")

	if guess.Fails != 0 {
		t.Errorf("0 failures, but it was %v", guess.Fails)
	}

	secondGuess, _ := h.Guess("p")

	if secondGuess.Fails != 1 {
		t.Errorf("1 failure , but it was %v", secondGuess.Fails)
	}

	thirdGuess, _ := h.Guess("p")

	if thirdGuess.Fails != 2 {
		t.Errorf("2 failures, but it was %v", thirdGuess.Fails)
	}
}

func TestGuessShouldRegisterTheLetterAttempted(t *testing.T) {
	h := Hangman{}
	h.SetWord("hola")

	guess, _ := h.Guess("h")
	if len(guess.LettersAttempted) != 1 && guess.LettersAttempted[0] != 'h' {
		t.Errorf("One letter was tried. Expected h, it was %v", guess.LettersAttempted[0])
	}

	secondGuess, _ := h.Guess("i")
	if len(secondGuess.LettersAttempted) != 2 && secondGuess.LettersAttempted[1] != 'i' {
		t.Errorf("Two letters were tried. Expected as second i, it was %v", secondGuess.LettersAttempted[0])
	}
}

func TestGuessTheSameLetterCanBeUsedOnlyOnce(t *testing.T) {
	h := Hangman{}
	h.SetWord("hola")

	guess, _ := h.Guess("h")
	if len(guess.LettersAttempted) != 1 && guess.LettersAttempted[0] != 'h' {
		t.Errorf("One letter was tried. Expected h, it was %v", guess.LettersAttempted[0])
	}

	secondGuess, _ := h.Guess("h")
	if len(secondGuess.LettersAttempted) != 1 {
		t.Errorf("Same letter was tried twice, should be recorded once.")
	}
}

func TestGuess(t *testing.T) {
	h := Hangman{}
	h.SetWord("puppa")
	guess, _ := h.Guess("p")

	assertWordIs(guess.WordGuessedSoFar, "p_pp_", t)
	assertWordIsFullyDiscovered(guess.IsGuessed, false, guess.WordGuessedSoFar, t)

	secondGuess, _ := h.Guess("u")

	assertWordIs(secondGuess.WordGuessedSoFar, "pupp_", t)
	assertWordIsFullyDiscovered(secondGuess.IsGuessed, false, secondGuess.WordGuessedSoFar, t)

	lastGuess, _ := h.Guess("a")

	assertWordIs(lastGuess.WordGuessedSoFar, "puppa", t)
	assertWordIsFullyDiscovered(lastGuess.IsGuessed, true, lastGuess.WordGuessedSoFar, t)
}

func TestGuessDifficultWord(t *testing.T) {
	h := Hangman{}
	h.SetWord("puppa mela")

	guess, _ := h.Guess("p")

	assertWordIs(guess.WordGuessedSoFar, "p_pp_ ____", t)
	assertWordIsFullyDiscovered(guess.IsGuessed, false, guess.WordGuessedSoFar, t)

	secondGuess, _ := h.Guess("u")

	assertWordIs(secondGuess.WordGuessedSoFar, "pupp_ ____", t)
	assertWordIsFullyDiscovered(secondGuess.IsGuessed, false, secondGuess.WordGuessedSoFar, t)

	thirdGuess, _ := h.Guess("m")

	assertWordIs(thirdGuess.WordGuessedSoFar, "pupp_ m___", t)
	assertWordIsFullyDiscovered(thirdGuess.IsGuessed, false, thirdGuess.WordGuessedSoFar, t)

	fourthGuess, _ := h.Guess("a")

	assertWordIs(fourthGuess.WordGuessedSoFar, "puppa m__a", t)
	assertWordIsFullyDiscovered(fourthGuess.IsGuessed, false, fourthGuess.WordGuessedSoFar, t)

	fifthGuess, _ := h.Guess("e")

	assertWordIs(fifthGuess.WordGuessedSoFar, "puppa me_a", t)
	assertWordIsFullyDiscovered(fifthGuess.IsGuessed, false, fifthGuess.WordGuessedSoFar, t)

	lastGuess, _ := h.Guess("l")

	assertWordIs(lastGuess.WordGuessedSoFar, "puppa mela", t)
	assertWordIsFullyDiscovered(lastGuess.IsGuessed, true, lastGuess.WordGuessedSoFar, t)
}

// helpers

func assertWordIs(encrypted string, word string, t *testing.T) {
	if encrypted != word {
		t.Errorf("Word was supposed to be %v, it was %v", word, encrypted)
	}
}

func assertWordIsFullyDiscovered(status bool, isCompleted bool, word string, t *testing.T) {
	if status != isCompleted {
		t.Errorf("I expected %v, I got %v. The word is %v", isCompleted, status, word)
	}
}
