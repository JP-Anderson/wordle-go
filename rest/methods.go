package rest

import (
	"fmt"
	"net/http"
	"wordle/engine"
	"wordle/rest/model"
	"github.com/gin-gonic/gin"
)

var games = make(map[string]*engine.Game)

//todo: move this to engine somewhere?
const defaultGuesses = 5

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/game", postGame)
	r.GET("/health", getHealth)
	return r
}

func postGame(c *gin.Context) {
	var newGame model.Game
	
	if err := c.BindJSON(&newGame); err != nil {
		return
	}

	id := newGame.UserID
	if _, ok := games[id]; ok {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprintf("game exists for user %s", id))
		return
	}
	
	// todo: add word list to this module and get random word
	createdGame := engine.New("snack", defaultGuesses)
	games[id] = createdGame
	newGame.TotalGuesses = defaultGuesses
	newGame.GameState = int(createdGame.Status)
	c.IndentedJSON(http.StatusOK, newGame)
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "wordle ok")
}
