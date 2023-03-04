package wordle

import (
	"math/rand"
	"time"
)

// TODO: import more words from dict txt
var words = []string{ "stork", "fight", "bough" }

func init() {
	rand.Seed(time.Now().Unix())
}

type Game struct {
	target string
	letterDic map[rune]bool
	guesses int
	status GameStatus
}

func New(target string, guesses int) *Game {
	dict := map[rune]bool{}
	for _, r := range target {
		dict[r] = true
	}
	return  &Game{
		target: target,
		letterDic: dict,
		guesses: guesses,
		status: GameInProgress,
	}
}

func (g *Game) IsFinished() bool {
	return g.status != GameInProgress
}

func (g *Game) guess(guess string) []letterStatus {
	if g.status != GameInProgress {
		return nil
	}
	g.guesses -= 1
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
	if g.target == guess {
		g.status = GameWon
		return result
	}
	if g.guesses == 0 {
		g.status = GameLost
	}
	return result
}

type GameStatus int64

const (
	GameInProgress GameStatus = 0
	GameLost GameStatus = 1
	GameWon GameStatus = 2
)

type letterStatus int64

const (
	LetterAndPositionCorrect letterStatus = 1
	LetterCorrectPositionIncorrect letterStatus = 0
	LetterIncorrect letterStatus = -1
)
