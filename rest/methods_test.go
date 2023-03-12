package rest

import (
	"encoding/json"
	"bytes"
	"fmt"	
	"net/http"
	"net/http/httptest"
	"testing"

	"wordle/rest/model"

	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func stubNextWordleWordFunc(target string) {
	NewWord = func() string {
		return target
	}
}

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
	stubNextWordleWordFunc("snack")

	w := httptest.NewRecorder()
	req := newGameRequest(t, "1")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	returnModel := responseRecorderToGameModel(t, w)
	assert.Equal(t, []*model.Guess{ nil, nil, nil, nil, nil }, returnModel.Guesses)
	assert.Equal(t, 0, returnModel.GameState)
	assert.Equal(t, 5, returnModel.TargetLength)
	assert.Nil(t, returnModel.Answer)

	t.Run("AndGetGameReturnsIdentical", func (t *testing.T) {
		getGameEndpointReturnsExpectedModel(t, "1", router, returnModel)
	})
}

func TestPostGameReturnsErrorWhenGameExistsForUserID(t *testing.T) {
	router := Router()
	stubNextWordleWordFunc("snack")

	w := httptest.NewRecorder()
	req := newGameRequest(t, "1")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	w2 := httptest.NewRecorder()	
	req2 := newGameRequest(t, "1")
	router.ServeHTTP(w2, req2)
	assert.Equal(t, 400, w2.Code)
	assert.Equal(t, "\"game exists for user 1\"", w2.Body.String())
}

func TestGetGameReturns404WithNoGameForUserID(t *testing.T) {
	router := Router()
	stubNextWordleWordFunc("snack")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/game/missing-user", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "\"no game exists for user missing-user\"", w.Body.String())
}

func TestPostGuessReturnsErrorWithNoGameForUserID(t *testing.T) {
	router := Router()
	stubNextWordleWordFunc("snack")

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
	stubNextWordleWordFunc("snack")

	w := httptest.NewRecorder()
	req := newGameRequest(t, "1")
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
	returnModel := responseRecorderToGameModel(t, w2)
	assert.Equal(t, "1", returnModel.UserID)
	guessModel := &model.Guess{
		GuessWord: "crane",
		LetterStatuses: "10210",
	}
	guesses := []*model.Guess{guessModel, nil, nil, nil, nil}
	assert.Equal(t, guesses, returnModel.Guesses)
	assert.Equal(t, 0, returnModel.GameState)
	assert.Equal(t, 5, returnModel.TargetLength)
	assert.Nil(t, returnModel.Answer)

	t.Run("AndGetGameReturnsIdentical", func (t *testing.T) {
		getGameEndpointReturnsExpectedModel(t, "1", router, returnModel)
	})
}

func TestGameVictory(t *testing.T) {
	router := Router()
	stubNextWordleWordFunc("snack")

	w := httptest.NewRecorder()
	req := newGameRequest(t, "1")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	
	w2 := httptest.NewRecorder()
	guessRequest := &model.GuessRequest{
		UserID: "1",
		Guess: "snack",
	}

	req2, _ := http.NewRequest("POST", "/guess", guessModelToBytesBuffer(t, guessRequest))
	router.ServeHTTP(w2, req2)
	assert.Equal(t, 200, w2.Code)
	returnModel := responseRecorderToGameModel(t, w2)
	assert.Equal(t, "1", returnModel.UserID)
	guessModel := &model.Guess{
		GuessWord: "snack",
		LetterStatuses: "22222",
	}
	guesses := []*model.Guess{guessModel, nil, nil, nil, nil}
	assert.Equal(t, guesses, returnModel.Guesses)
	assert.Equal(t, 2, returnModel.GameState)
	assert.Equal(t, 5, returnModel.TargetLength)
	assert.Equal(t, "snack", returnModel.Answer.Answer)
	t.Run("AndGetGameReturnsIdentical", func (t *testing.T) {
		getGameEndpointReturnsExpectedModel(t, "1", router, returnModel)
	})
}

func newGameRequest(t *testing.T, id string) *http.Request {
	newGameRequest := &model.Game{
		UserID: "1",
	}
	req, _ := http.NewRequest("POST", "/game", gameModelToBytesBuffer(t, newGameRequest))
	return req
}

func getGameEndpointReturnsExpectedModel(t *testing.T, id string, router *gin.Engine, expectedReturnModel *model.Game) {
	getGame, _ := http.NewRequest("GET", fmt.Sprintf("/game/%s", id), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, getGame)
	assert.Equal(t, 200, w.Code)
	getReturnModel := responseRecorderToGameModel(t, w)
	assert.Equal(t, expectedReturnModel, getReturnModel)
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

