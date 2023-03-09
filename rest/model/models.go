package model

// Guess represents a guess which has been made in the past.
// - GuessWord is the actual guess made by the player.
// - LetterStatuses is a string the same length as the guess word, each index contains a number indicating the correctness of the same index in the guess.
type Guess struct {
	GuessWord string `json:"guess_word"`
	LetterStatuses string `json:"letter_statuses"`
}

// GuessRequest is the JSON Request Body sent to the POST /guess endpoint
type GuessRequest struct {
	UserID string `json:"user_id"`
	Guess string `json:"guess"`	
}

// Game represents a game of Wordle which may be complete or in progress.
// This will be returned from the GET/game and POST/guess endpoints.
type Game struct {
	UserID string `json:"user_id"`
	Guesses []Guess `json:"guesses"`
	TotalGuesses int `json:"total_guesses"`
	GameState int `json:"game_state"`
}

// Answer contains the answer for a selected game.
// This has its own endpoint because we do not return the Answer in the regular Game state structure.
// This will only be returned for completed (won or lost) games.
type Answer struct {
	Answer string `json:"answer"`
}
