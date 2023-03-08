package main

import (
	"fmt"
	"wordle/engine"
	"wordle/words"
)

type WordleSignal int64

const (
	Start WordleSignal = iota
	MakeNextGuess
	Finished
)

func main() {
	wl := words.WordsListFromFile("../words/dic.txt")
	guesses := 5
	g := engine.New(wl.NextWord(), guesses)
	guessChan := make(chan string)
	signalChan := make(chan WordleSignal)
	go run(g, guessChan, signalChan)
	for {
		sig := <- signalChan
		if sig == Finished {
			if g.Status == engine.GameWon {
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

func run(g *engine.Game, guessChan <-chan string, signalChan chan<- WordleSignal) {
	signalChan <- Start
	for !g.IsFinished() {
		guess := <- guessChan
		if len(guess) != 5 {
			continue
		}
		result := g.Guess(guess)
		fmt.Println(result)
		if g.IsFinished() {
			signalChan <- Finished
		}
		signalChan <- MakeNextGuess
	}
	fmt.Println("Game result ", g.Status)
}
