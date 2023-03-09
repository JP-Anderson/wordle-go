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
		GameID: "1",
		UserID: "1",		
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(newGameRequest)
	req, _ := http.NewRequest("POST", "/game", payloadBuf)
	
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var returnModel model.Game
	err := json.Unmarshal(w.Body.Bytes(), &returnModel)
	assert.NoError(t, err)
	assert.Equal(t, 0, returnModel.GameState)
	assert.Equal(t, 5, returnModel.TotalGuesses)
}
