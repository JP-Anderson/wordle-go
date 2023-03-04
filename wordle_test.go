package wordle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWordlePicksTargetWord(t *testing.T) {
	game := New("shark")
	assert.Equal(t, "shark", game.target)
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
		assert.Equal(t, tc.output, New(tc.target).guess(tc.guess))
	}
}
