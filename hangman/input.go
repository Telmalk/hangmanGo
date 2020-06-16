package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)


func isLetter(c string) bool {
	return ("a" <= c && c <= "z") || ("A" <= c && c <= "Z")
}

func isNum(c string) bool {
	return "1" <= c && c <= "3"
}


func ReadGuess() (guess string, err error)  {
	valid := false
	for !valid {
		fmt.Print("What is your letter ? or tap hint to get a hint ")
		guess, err = reader.ReadString('\n')
		if err != nil {
			return guess, err
		}
		guess = strings.TrimSpace(guess)
		if guess == "hint" {
			return guess, nil
		} else if len(guess) != 1 || !isLetter(guess)  {
			fmt.Printf("Invalid letter input. letter=%v, len=%v\n", guess, len(guess))
			continue
		}
		valid = true
	}
	return guess, nil
}

func ChooseDificult() (difficulty string, err error) {
	valid := false
	for !valid {
		fmt.Print("What is your difficulty ? ")
		difficulty, err = reader.ReadString('\n')
		if err != nil {
			return difficulty, err
		}
		difficulty = strings.TrimSpace(difficulty)
		if   len(difficulty) != 1 || !isNum(difficulty)  {
			fmt.Printf("Expect 1 or 2 or 3 got: %v\n", difficulty)
			continue
		}
		valid = true
	}
	return difficulty, nil

}
