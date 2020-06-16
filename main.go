package main

import (
	"fmt"
	"os"
	"training.go/hangman/dictionary"
	"training.go/hangman/hangman"
)


func main()  {
	err := dictionary.Load("word.txt")
	if err != nil {
		fmt.Printf("Could not load dictionary: %v", err)
		os.Exit(1)
	}
	hangman.DrawDificult()
	diff, err := hangman.ChooseDificult()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	g, err := hangman.New(8, dictionary.PicWord(), hangman.GetNbHint(diff))
	if err != nil {
		fmt.Printf("Err: %v", err)
	}
	hangman.DrawWelcomde()

	guess := ""
	i := 1
	for {
		hangman.Draw(g, guess, i)
		switch g.State {
		case "won", "lost":
			os.Exit(0)

		}
		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal %v\n", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
		i++
	}
}