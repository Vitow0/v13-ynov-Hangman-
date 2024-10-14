package functions

import (
	"strings"
)

// Initialize the game
func InitializeGame(word string, hardMode bool) []rune {
	wordRunes := []rune(word) // Convert word in runes

	var revealedLetters int
	if hardMode {
		revealedLetters = len(wordRunes)/3 - 1 // reveal letter in hard mode
	} else {
		revealedLetters = len(wordRunes)/2 - 1 // reveal letter in normal mode
	}

	hiddenWordRunes := make([]rune, len(wordRunes))
	for i := 0; i < len(wordRunes); i++ {
		if i < revealedLetters {
			hiddenWordRunes[i] = wordRunes[i] // render the first letter
		} else {
			hiddenWordRunes[i] = '_'
		}
	}

	return hiddenWordRunes // return hiddenword
}

// Submit a letter
func submitLetter(guessRune rune, normalizedWord []rune, hiddenWord *[]rune, hardMode bool, vowelCount *int, stageOfDeath *int) string {

	vowels := "AEOUIY"      // list of vowel
	correctGuess := false   // indicate if guess is correct
	penaltyApplied := false // indicate if penality is applied
	message := ""           // message to return

	// Check if the letter is a vowel and update the hidden word
	for i, r := range normalizedWord {
		if strings.EqualFold(string(r), string(guessRune)) && (*hiddenWord)[i] == '_' {
			(*hiddenWord)[i] = r
			correctGuess = true
		} else {
			// Call NormalizeRune to check for accented characters
			if guessRune == normalizeLetter(r) && (*hiddenWord)[i] == '_' {
				(*hiddenWord)[i] = r
				correctGuess = true
			}
		}
	}

	if strings.ContainsRune(vowels, guessRune) { // Increment vowel count if the guess is a vowel (regardless if it's correct or not)
		*vowelCount++
	}

	// Apply penalty if the vowel is correct and the player uses 3 vowels or more
	if correctGuess && strings.ContainsRune(vowels, guessRune) && hardMode && *vowelCount >= 3 && !penaltyApplied {
		message = "Pénalité appliquée pour avoir utilisé plus de 3 voyelles."
		*stageOfDeath++
		penaltyApplied = true
	}

	if !correctGuess && !penaltyApplied { // Apply penalty if the letter is not correct (or if already revealed)
		*stageOfDeath++
		message = "Lettre incorrecte."
	}

	return message // return the message for these condition
}

// Function to add the variants letters
func normalizeLetter(r rune) rune {
	switch r {
	case 'É', 'È', 'Ê', 'Ë':
		return 'E'
	case 'Á', 'À', 'Â', 'Ä':
		return 'A'
	case 'Í', 'Ì', 'Î', 'Ï':
		return 'I'
	case 'Ó', 'Ò', 'Ô', 'Ö':
		return 'O'
	case 'Ú', 'Ù', 'Û', 'Ü':
		return 'U'
	case 'Ç':
		return 'C'
	default:
		return r // return if there is no accent
	}
}
