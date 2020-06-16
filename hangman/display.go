package hangman

import "fmt"

func DrawWelcomde()  {
	fmt.Println(`
	|\     /|(  ___  )( (    /|(  ____ \(       )(  ___  )( (    /|
	| )   ( || (   ) ||  \  ( || (    \/| () () || (   ) ||  \  ( |
	| (___) || (___) ||   \ | || |      | || || || (___) ||   \ | |
	|  ___  ||  ___  || (\ \) || | ____ | |(_)| ||  ___  || (\ \) |
	| (   ) || (   ) || | \   || | \_  )| |   | || (   ) || | \   |
	| )   ( || )   ( || )  \  || (___) || )   ( || )   ( || )  \  |
	|/     \||/     \||/    )_)(_______)|/     \||/     \||/    )_)
	`)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |   / \
  _|_
 |   |______
 |          |
 |__________|
		`
	case 1:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |    
  _|_
 |   |______
 |          |
 |__________|
		`
	case 2:
		draw = `
    ____
   |    |      
   |    o      
   |    |
   |    |
   |     
  _|_
 |   |______
 |          |
 |__________|
		`
	case 3:
		draw = `
    ____
   |    |      
   |    o      
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 4:
		draw = `
    ____
   |    |      
   |      
   |      
   |  
   |  
  _|_
 |   |______
 |          |
 |__________|
		`
	case 5:
		draw = `
    ____
   |        
   |        
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 6:
		draw = `
    
   |     
   |     
   |     
   |
   |
  _|_
 |   |______
 |          |
 |__________|
		`
	case 7:
		draw = `
  _ _
 |   |______
 |          |
 |__________|
		`
	case 8:
		draw = `

		`
	}
	fmt.Println(draw)
}

func drawLetters(l [] string)  {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}

func DrawDificult()  {
	fmt.Println("Chose '1', '2', or '3'")
	fmt.Println("1: easy 3 indices")
	fmt.Println("2: normal 1 indices")
	fmt.Println("3: hard 0 indices")
}

func drawState(g *Game, guess string)  {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "googGuess":
		fmt.Print("Good guess !")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used", guess)
	case "badGuess":
		fmt.Printf("Bad guess, '%s' is not in the word", guess)
	case "hint":
		fmt.Printf("The word contains letter(s) '%v' ", g.LetterHint)
	case "notEnoughtHint":
		fmt.Printf("You don't have hint in stock stock: %v", g.Hint)
	case "lost":
		fmt.Print("You lost :( The wordWard: ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("You Won the word was: ")
		drawLetters(g.Letters)
	}

}

func infoGame(turnLeft, nbHint, turn int, letterHint []string)  {
	fmt.Printf("Tour: %v\n", turn)
	fmt.Printf("Vie restante: %v\n", turnLeft)
	fmt.Printf("Indices retant: %v\n", nbHint)
	fmt.Printf("Hint: %v\n", letterHint)
	fmt.Println("Tap hint to get a hint")
}

func Draw(g *Game, guess string, i int)  {
	if i > 1 {
		print("\033[H\033[2J")
	}
	infoGame(g.TurnsLeft, g.Hint, i, g.LetterHint)
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}
