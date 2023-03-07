package main

import (
	"fmt"
	"wordle"
)

func main() {
	g := wordle.New("puppy", 5)
	finChan := make(chan bool)
	guessChan := make(chan string)
	go run(g, guessChan, finChan)
	loop:
	for {
		fmt.Println("for->")
		select {
		case _ = <- finChan:
			break loop
		default:
			fmt.Println("Guess next word")
			var guess string
			fmt.Scan(&guess)
			if len(guess) != 5 {
				fmt.Println("Must be 5 letter word")
				continue
			}
			guessChan<-guess
			
		}
	}
	close(guessChan)
	close(finChan)
	fmt.Println("get here")
}

func run(g *wordle.Game, guessChan <-chan string, finChan chan<- bool ) {
	for !g.IsFinished() {
		guess := <- guessChan
		if len(guess) != 5 {
			continue
		}
		result := g.Guess(guess)
		fmt.Println(result)
	}
	fmt.Println("Game result ", g.Status)
	finChan <- true
	fmt.Println("finished run")
}
