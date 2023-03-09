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

	var returnModel model.Game
	err := json.Unmarshal(w.Body.Bytes(), &returnModel)
	assert.NoError(t, err)
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

func gameModelToBytesBuffer(t *testing.T, game *model.Game) *bytes.Buffer {
	buf := new(bytes.Buffer)
	assert.NoError(t, json.NewEncoder(buf).Encode(game))
	return buf
}

