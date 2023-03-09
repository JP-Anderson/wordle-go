package rest

import (
	"fmt"
	"net/http"
	"wordle/engine"
	"wordle/rest/model"
	"github.com/gin-gonic/gin"
)

var games map[string]*engine.Game

//todo: move this to engine somewhere?
const defaultGuesses = 5

func Router() *gin.Engine {
	games = make(map[string]*engine.Game)
	r := gin.Default()
	r.POST("/game", postGame)
	r.POST("/guess", postGuess)
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
	newGame.GameState = int(createdGame.Status)
	newGame.Guesses = createdGame.Guesses()
	c.IndentedJSON(http.StatusOK, newGame)
}

func postGuess(c *gin.Context) {
	var newGuess model.GuessRequest
	
	if err := c.BindJSON(&newGuess); err !=  nil {
		return
	}

	id := newGuess.UserID
	var game *engine.Game
	var ok bool
	if game, ok = games[id]; !ok {
		c.IndentedJSON(http.StatusNotFound, fmt.Sprintf("game does not exist for user %s", id))
		return
	}

	// TODO check game is finished first (will not do this until guess functionality is finished so we can test game is finished)
	game.Guess(newGuess.Guess)
	
	returnModel := &model.Game{
		Guesses: game.Guesses(),
		UserID: id,
		GameState: int(game.Status),
	}
	c.IndentedJSON(http.StatusOK, returnModel)
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "wordle ok")
}
