package main

import (
	"fmt"
	"wordle"
	"wordle/words"
)

func main() {
	wl := words.WordsListFromFile("../words/dic.txt")
	guesses := 5
	g := wordle.New(wl.NextWord(), guesses)
	guessChan := make(chan string)
	signalChan := make(chan wordle.GameSignal)
	go g.Start(guessChan, signalChan)
	for {
		sig := <- signalChan
		if sig == wordle.Finished {
			if g.Status == wordle.GameWon {
				fmt.Printf("You won with %d guesses remaining!\n", guesses)
			} else {
				fmt.Printf("Unlucky, the word was %s\n", g.Target())
			}
			break
		}
		var guess string
		inputLoop:
		for {
			fmt.Printf("Guess next word. %d guesses left!\n", guesses)
			fmt.Scan(&guess)
			if len(guess) != 5 {
				fmt.Println("Must be 5 letter word")
				continue inputLoop
			}
			if !wl.Valid(guess) {
				fmt.Printf("%s is not in the dictionary\n", guess)
				continue inputLoop
			}
			break
		}
		guesses = guesses-1
		guessChan<-guess
	}
	close(guessChan)
}
