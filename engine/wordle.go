package engine

import (
	"wordle/rest/model"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Game struct {
	target string
	letterDic map[rune]bool
	guesses []*model.Guess
	Status GameStatus
}

func New(target string, guesses int) *Game {
	dict := map[rune]bool{}
	for _, r := range target {
		dict[r] = true
	}
	return  &Game{
		target: target,
		letterDic: dict,
		guesses: make([]*model.Guess, guesses),
		Status: GameInProgress,
	}
}

func (g *Game) GuessesMade() int {
	guessesMade := 0
	for _, guess := range g.guesses {
		if guess == nil {
			break
		}
		guessesMade += 1
	}
	return guessesMade
}

func (g *Game) GuessesRemaining() int {
	return len(g.guesses) - g.GuessesMade()
}

func (g *Game) IsFinished() bool {
	return g.Status != GameInProgress
}

func (g *Game) Guess(guess string) []letterStatus {
	if g.Status != GameInProgress {
		return nil
	}
	result := make([]letterStatus, len(g.target))
	for i, ch := range g.target {
		if ch == rune(guess[i]) {
			result[i] = LetterAndPositionCorrect
		} else if _, ok := g.letterDic[rune(guess[i])]; ok {
			result[i] = LetterCorrectPositionIncorrect
		} else {
			result[i] = LetterIncorrect
		}
	}
	nextGuessIx := g.GuessesMade()
	guessModel := &model.Guess{
		GuessWord: guess,
		LetterStatuses: letterStatusesToString(result),	
	}
	g.guesses[nextGuessIx] = guessModel
	if g.target == guess {
		g.Status = GameWon
		return result
	}
	if nextGuessIx == len(g.guesses)-1 {
		g.Status = GameLost
	}
	return result
}

func (g *Game) Target() string {
	return g.target
}

func (g *Game) ToApiModel(userID string) *model.Game {
	game := &model.Game{
		UserID: userID,
		Guesses: g.guesses,
		GameState: int(g.Status),
	}
	if g.IsFinished() {
		game.Answer = &model.Answer{
			Answer: g.target,	
		}
	}
	return game
}

func letterStatusesToString(statuses []letterStatus) string {
	sb := strings.Builder{}
	for _, s := range statuses {
		sb.WriteString(fmt.Sprintf("%d", s))
	}
	return sb.String()
}

type GameStatus int64

const (
	GameInProgress GameStatus = 0
	GameLost GameStatus = 1
	GameWon GameStatus = 2
)

type letterStatus int64

const (
	LetterAndPositionCorrect letterStatus = 2
	LetterCorrectPositionIncorrect letterStatus = 1
	LetterIncorrect letterStatus = 0
)
