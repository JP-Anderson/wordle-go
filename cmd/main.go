package main

import (
	"fmt"
	"wordle"
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
	g := wordle.New(wl.NextWord(), guesses)
	guessChan := make(chan string)
	signalChan := make(chan WordleSignal)
	go run(g, guessChan, signalChan)
	for {
		sig := <- signalChan
		if sig == Finished {
			if g.Status == wordle.GameWon {
				fmt.Printf("You won with %d guesses remaining!\n", guesses)
			} else {
				fmt.Printf("Unlucky, the word was %s\n", g.Target())
			}
			break
		}
		fmt.Printf("Guess next word. %d guesses left!\n", guesses)
		var guess string
		fmt.Scan(&guess)
		if len(guess) != 5 {
			fmt.Println("Must be 5 letter word")
			continue
		}
		guesses = guesses-1
		guessChan<-guess
	}
	close(guessChan)
}

func run(g *wordle.Game, guessChan <-chan string, signalChan chan<- WordleSignal) {
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
