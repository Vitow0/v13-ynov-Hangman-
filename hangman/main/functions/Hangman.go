package functions

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode/utf8"
)

// Select a word randomly
func selectRandomWord(mot string) (string, error) {
	file, err := os.Open(mot) // open file
	if err != nil {
		return "Fichier non trouvé ou manquant", err
	}
	defer file.Close() // check if file is close in the end of function

	words := []string{}
	scanner := bufio.NewScanner(file) // add a scanner
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	random := rand.Intn(len(words)) // choose a random word
	return words[random], nil
}

// Play the game
func PlayHangman(hardMode bool) {
	menu() // call menu

	for {
		fmt.Print("\n\n\n\n\n\n\n\n")

		word, err := selectRandomWord("mot.txt") // call selectRandomWord
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier:", err) // end if there is an error
			return
		}

		hiddenWord := []rune(InitializeGame(word, hardMode))
		normalizedWord := []rune(word)     // convert the word in runes for the accents
		vowelCount := 0                    // number of vowel
		stageOfDeath := 0                  // case of hangman
		lifes := 10                        // number of life
		guessedLetters := ""               // letter that has been guessed
		usedLetters := make(map[rune]bool) // letter that has been used
		var message string                 // display the message to the player

		for stageOfDeath < lifes && string(hiddenWord) != string(normalizedWord) { // render of the game
			drawHangman(stageOfDeath)
			fmt.Println("Tapez 'exit' si vous voulez quitter")
			fmt.Println("Mot à deviner:", string(hiddenWord))
			if !hardMode {
				fmt.Println("Lettres devinées:", guessedLetters)
			}
			fmt.Print("Choisissez une lettre ou devinez le mot (en mettant les accents si besoin): ")
			fmt.Printf("Il vous reste %d tentatives.\n", lifes-stageOfDeath)

			if message != "" { // display message for each condition
				fmt.Println(message)
				message = ""
			}

			var guess string
			fmt.Scanln(&guess) // read the imput of user

			guess = strings.ToUpper(guess) // convert letter in  capital letter

			if strings.ToLower(guess) == "exit" { // leave the game
				fmt.Println("Vous avez quitté le jeu.")
				return
			}

			if len(guess) > 1 {
				if guess == word { // if the word is correct
					hiddenWord = []rune(word)
				} else { //else remove 2 lifes
					message = "Mauvais Mot ! Vous perdez 2 vies."
					stageOfDeath += 2
				}
				continue
			}

			guessRune, _ := utf8.DecodeRuneInString(guess)    // decode the hidden word in rune
			normalizedGuessRune := normalizeLetter(guessRune) // Normalize rune for the accents

			if usedLetters[normalizedGuessRune] { // if letter is already used then display this message in the select mode
				message = "Lettre déjà utilisée."
				if hardMode {
					message = "Vous perdez une vie pour avoir réutilisé une lettre."
					stageOfDeath++
				}
				continue
			}

			usedLetters[normalizedGuessRune] = true // display the guess letter at the list in normal mode
			if !hardMode {
				guessedLetters += guess
			}

			// call the function submit letter
			message = submitLetter(normalizedGuessRune, normalizedWord, &hiddenWord, hardMode, &vowelCount, &stageOfDeath)

		}
		// if you win or lose, then it display this message
		if string(hiddenWord) == string(normalizedWord) {
			fmt.Println("Félicitations ! Vous avez deviné le mot :", string(normalizedWord))
		} else {
			drawHangman(lifes) // final case of hangman
			fmt.Println("Vous avez perdu ! Le mot était :", string(normalizedWord))
		}

		var replay string
		fmt.Print("Appuyez sur 'Entrée' pour rejouer ou tapez 'exit' pour quitter: ")
		fmt.Scanln(&replay)                    // replay the game
		if strings.ToLower(replay) == "exit" { // leave the game
			fmt.Println("Merci d'avoir joué !")
			break
		}
	}
}

// function to display the hangman
func drawHangman(stage_of_death int) {

	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n") // space between each guess

	if stage_of_death == 0 { // it's for not display the first case of hangman in the file
		return
	}

	file, err := os.Open("hangman.txt")
	if err != nil { // if the file is not found, then display error
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	defer file.Close() // ensure file is close at the end of function

	hangman := []string{}
	scanner := bufio.NewScanner(file)

	// Scan the file
	for scanner.Scan() {
		hangman = append(hangman, scanner.Text()) // display the text
	}

	// Display the case of hangman
	stageSize := 8           // each case has 8 lines
	if stage_of_death == 0 { // Don't print for the initial case
		return
	}

	// calcul for display each case
	start := (stage_of_death - 1) * stageSize
	end := start + stageSize

	if start < len(hangman) { // index for the lenght of the hangman
		if end > len(hangman) {
			end = len(hangman)
		}

		for i := start; i < end; i++ {
			fmt.Println(hangman[i]) // Display the lines for the current hangman stage
		}
	}
}

// function to display the menu
func menu() {
	// add colors in the menu
	red := "\033[31m"
	green := "\033[32m"
	yellow := "\033[33m"
	reset := "\033[0m"

	// render the menu
	fmt.Println(green + `  _    _                                             |` + reset)
	fmt.Println(green + ` | |  | |                                            |` + reset)
	fmt.Println(green + ` | |__| | __ _ _ __   ___  _ ___ ___  ___  _ __      |   __  __  _ __    ______  __     __ ` + reset)
	fmt.Println(yellow + ` |  __  |/ _\` + "`" + `| '_ \ / _ \` + `| '_  '_` + `  |` + `/` + "`" + ` _\| '_ \     |   \ \/ / | '_ \  |  __  | \ \   / /` + reset)
	fmt.Println(yellow + ` | |  | | (_| | | | | (_| | | | | | | (_| | | | |    |    \  /  | | | | | |__| |  \ \_/ /` + reset)
	fmt.Println(red + ` |_|  |_|\__,_|_| |_|\__, |_| |_| |_|\__,_|_| |_|    |    /_/   |_| |_| |______|   \___/ ` + reset)
	fmt.Println(red + `                      __/ |                          |   ` + reset)
	fmt.Println(red + `                     |___/                           |   ` + reset)

	fmt.Println("Appuyez sur 'Entrée' pour commencer le jeu ou tapez 'exit' pour quitter")

	reader := bufio.NewReader(os.Stdin) // begin the game
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)	

	if strings.ToLower(input) == "exit" { // leave in the end of the game
		fmt.Println("Vous avez quitté le jeu.")
		os.Exit(0)
	}
}
