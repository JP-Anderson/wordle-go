package wordle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const defaultGuesses = 5

func TestNewWordlePicksTargetWord(t *testing.T) {
	game := New("shark", defaultGuesses)
	assert.Equal(t, "shark", game.target)
}

func TestFailingAGame(t *testing.T) {
	game := New("doggy", defaultGuesses)
	for i := 0; i < defaultGuesses-1; i++ {
		game.guess("catty")
		assert.False(t, game.IsFinished())
		assert.Equal(t, GameInProgress, game.status)
	}
	game.guess("catty")
	assert.True(t, game.IsFinished())
	assert.Equal(t, GameLost, game.status)
	
	// further guesses should return nil for a finished game.
	result := game.guess("catty")
	assert.Nil(t, result)
}

func TestGuess(t *testing.T) {
	type testCase struct {
		title string
		target string
		guess string
		output []letterStatus
	}

	cases := []testCase{
		{
			"correct letter",
			"crane",
			"clock",
			[]letterStatus{LetterAndPositionCorrect, LetterIncorrect, LetterIncorrect, LetterCorrectPositionIncorrect, LetterIncorrect},
		},
		{
			"incorrect letter",
			"shark",
			"shirk",
			[]letterStatus{LetterAndPositionCorrect, LetterAndPositionCorrect, LetterIncorrect, LetterAndPositionCorrect, LetterAndPositionCorrect},	
		},
		{
			"letter correct but wrong place",
			"crane",
			"pluck",
			[]letterStatus{LetterIncorrect, LetterIncorrect, LetterIncorrect, LetterCorrectPositionIncorrect, LetterIncorrect},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.output, New(tc.target, defaultGuesses).guess(tc.guess))
	}
}
