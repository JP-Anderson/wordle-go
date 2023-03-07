package main

import (
	"fmt"
	"sync"
	"wordle"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	guesses := 5
	g := wordle.New("puppy", guesses)
	guessChan := make(chan string)
	go run(g, guessChan, wg)
	go func() {
		for {
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
	}()
	wg.Wait()
	close(guessChan)
}

func run(g *wordle.Game, guessChan <-chan string, wg *sync.WaitGroup) {
	for !g.IsFinished() {
		guess := <- guessChan
		if len(guess) != 5 {
			continue
		}
		result := g.Guess(guess)
		fmt.Println(result)
	}
	fmt.Println("Game result ", g.Status)
	wg.Done()
}
