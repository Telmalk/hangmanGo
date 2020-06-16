package hangman

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Game struct {
	State string
	Letters []string
	FoundLetters []string
	UsedLetters []string
	TurnsLeft int
	Hint int
	LetterHint []string

}

func New(turns int, word string, hint int) (*Game, error)  {
 	if len(word) < 2 {
 		return nil, fmt.Errorf("Word '%s' must be a least 2 characters. got %v ", word, len(word))
	}
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}
	g := &Game{
		State: "",
		Letters: letters,
		FoundLetters: found,
		UsedLetters: []string{},
		TurnsLeft: turns,
		Hint: hint,
	}
	return g, nil
}

func letterInWord(guess string, letters []string) bool  {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}

func (g *Game) RevealLetter(guess string)  {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}

func (g *Game) loseTurn(guess string)  {
	g.UsedLetters = append(g.UsedLetters, guess)
	g.TurnsLeft -= 1
}

func hasWon(letters []string, foundLetters[]string) bool  {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}

func (g *Game) GetLetter(guess string) bool {
	if g.Hint <= 0 || guess != "HINT"  {
		return false
	}
	g.Hint--
	rand.Seed(time.Now().Unix())
	g.LetterHint = append(g.LetterHint, g.Letters[rand.Intn(len(g.Letters))])
	return true
}

func GetNbHint(difficulty string) int  {
	switch difficulty {
	case "1":
		return 3
	case "2":
		return 1
	default:
		return 0
	}
}

func (g *Game) MakeAGuess(guess string)  {
	guess = strings.ToUpper(guess)
	switch g.State {
	case "won", "lost":
		return;
	}
	if g.Hint <= 0 && guess == "HINT" {
		g.State = "notEnoughtHint"
	} else if g.GetLetter(guess) {
		g.State = "hint"
	} else if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)
		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	} else   {
		g.State = "badGuess"
		g.loseTurn(guess)
		if g.TurnsLeft <= 0 {
			g.State = "lost"
		}
	}
}
