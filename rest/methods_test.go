package rest

import (
	"encoding/json"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"wordle/rest/model"

	"github.com/stretchr/testify/assert"
)

func TestHealthRoute(t *testing.T) {
	router := Router()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "\"wordle ok\"", w.Body.String())
}

func TestPostGameReturnsNewGame(t *testing.T) {
	router := Router()
	
	w := httptest.NewRecorder()
	newGameRequest := &model.Game{
		UserID: "1",		
	}
	req, _ := http.NewRequest("POST", "/game", gameModelToBytesBuffer(t, newGameRequest))
	
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	returnModel := responseRecorderToGameModel(t, w)
	assert.Equal(t, 0, returnModel.GameState)
	assert.Equal(t, 5, returnModel.TotalGuesses)
}

func TestPostGameReturnsErrorWhenGameExistsForUserID(t *testing.T) {
	router := Router()
	
	w := httptest.NewRecorder()
	newGameRequest := &model.Game{
		UserID: "1",
	}
	req, _ := http.NewRequest("POST", "/game", gameModelToBytesBuffer(t, newGameRequest))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	w2 := httptest.NewRecorder()	
	req2, _ := http.NewRequest("POST", "/game", gameModelToBytesBuffer(t, newGameRequest))	
	router.ServeHTTP(w2, req2)
	assert.Equal(t, 400, w2.Code)
	assert.Equal(t, "\"game exists for user 1\"", w2.Body.String())
}

func TestPostGuessReturnsErrorWithNoGameForUserID(t *testing.T) {
	router := Router()
	
	w := httptest.NewRecorder()
	guessRequest := &model.GuessRequest{
		UserID: "id-with-no-game",
		Guess: "crane",	
	}

	req, _ := http.NewRequest("POST", "/guess", guessModelToBytesBuffer(t, guessRequest))
	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "\"game does not exist for user id-with-no-game\"", w.Body.String())
}

func TestPostGuessReturnsGameStateWithGuessStatus(t *testing.T) {
	router := Router()

	w := httptest.NewRecorder()
	newGameRequest := &model.Game{
		UserID: "1",
	}
	req, _ := http.NewRequest("POST", "/game", gameModelToBytesBuffer(t, newGameRequest))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	w2 := httptest.NewRecorder()
	guessRequest := &model.GuessRequest{
		UserID: "1",
		Guess: "crane",
	}

	req2, _ := http.NewRequest("POST", "/guess", guessModelToBytesBuffer(t, guessRequest))
	router.ServeHTTP(w2, req2)
	assert.Equal(t, 200, w2.Code)
}

func gameModelToBytesBuffer(t *testing.T, game *model.Game) *bytes.Buffer {
	buf := new(bytes.Buffer)
	assert.NoError(t, json.NewEncoder(buf).Encode(game))
	return buf
}

func guessModelToBytesBuffer(t *testing.T, guess *model.GuessRequest) *bytes.Buffer {
	buf := new(bytes.Buffer)
	assert.NoError(t, json.NewEncoder(buf).Encode(guess))
	return buf
}

func responseRecorderToGameModel(t *testing.T, w *httptest.ResponseRecorder) *model.Game {
	var returnModel model.Game
	err := json.Unmarshal(w.Body.Bytes(), &returnModel)
	assert.NoError(t, err)
	return &returnModel
}

