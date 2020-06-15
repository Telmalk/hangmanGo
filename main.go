package main

import (
	"fmt"
	"os"
	"training.go/hangman/hangman"
	"training.go/hangman/hangman/dictionary"
)

func main()  {
	err := dictionary.Load("word.txt")
	if err != nil {
		fmt.Printf("Could not load dictionary: %v", err)
		os.Exit(1)
	}
	g := hangman.New(8, dictionary.PicWord())
	hangman.DrawWelcomde()

	guess := ""
	for {
		hangman.Draw(g, guess)

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
	}
}