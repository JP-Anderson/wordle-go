package engine

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
		game.Guess("catty")
		assert.False(t, game.IsFinished())
		assert.Equal(t, GameInProgress, game.Status)
		assert.Equal(t, i+1, game.GuessesMade())
	}
	assert.Equal(t, 4, game.GuessesMade())
	assert.Equal(t, 1, game.GuessesRemaining())
	game.Guess("catty")
	assert.True(t, game.IsFinished())
	assert.Equal(t, GameLost, game.Status)
	
	// further guesses should return nil for a finished game.
	result := game.Guess("catty")
	assert.Nil(t, result)
}

func TestWinningAGame(t *testing.T) {
	// Test that game can be won on turns 1 up to 5 (last guess).
	turns := []int{1,2,3,4,5}
	wrongGuesses := []string{"wrong", "guess", "small", "chief"}
	statusStrings := []string{"00000", "00222", "10000", "22010"}
	for _, turnNum := range turns {
		game := New("chess", defaultGuesses)
		for i := 0; i < turnNum-1; i++ {
			game.Guess(wrongGuesses[i])
			assert.Equal(t, wrongGuesses[i], game.guesses[i].GuessWord)
			assert.Equal(t, statusStrings[i], game.guesses[i].LetterStatuses)	
		}
		game.Guess("chess")
		assert.Equal(t, "chess", game.guesses[turnNum-1].GuessWord)
		assert.Equal(t, "22222", game.guesses[turnNum-1].LetterStatuses)
		assert.True(t, game.IsFinished())
		assert.Equal(t, GameWon, game.Status)
		assert.Nil(t, game.Guess("chess"))
	}
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
		g := New(tc.target, defaultGuesses)
		assert.Equal(t, tc.output, g.Guess(tc.guess))
		assert.Equal(t, tc.guess, g.guesses[0].GuessWord)
		assert.Equal(t, letterStatusesToString(tc.output), g.guesses[0].LetterStatuses)
	}
}

func TestLetterStatusesToString(t *testing.T) {
	type testCase struct {
		in []letterStatus
		out string
	}

	cases := []testCase{
		{
			[]letterStatus{LetterIncorrect, LetterIncorrect, LetterIncorrect, LetterIncorrect, LetterIncorrect },
			"00000",
		},
		{
			[]letterStatus{LetterIncorrect, LetterCorrectPositionIncorrect, LetterAndPositionCorrect, LetterAndPositionCorrect},
			"0122",
		},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.out, letterStatusesToString(tc.in))
	}
}
