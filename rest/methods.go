package rest

import (
	"fmt"
	"net/http"
	"strings"
	"wordle/engine"
	"wordle/rest/model"
	"wordle/rest/middleware"
	"wordle/words"
	"github.com/gin-gonic/gin"
)

var games map[string]*engine.Game
var wl words.WordsList

func Router() *gin.Engine {
	games = make(map[string]*engine.Game)
	wl = words.WordsListFromFile("words/dic.txt")
	r := gin.Default()
	r.Use(middleware.CORS())
	r.POST("/game", postGame)
	r.GET("/game/:id", getGame)
	r.POST("/guess", postGuess)
	r.GET("/health", getHealth)
	return r
}

var NewWord = func () string {
	return wl.NextWord()
}

var ValidWord = func (w string) bool {
	return wl.Valid(strings.ToLower(w))
}

func postGame(c *gin.Context) {
	var newGame model.Game
	
	if err := c.BindJSON(&newGame); err != nil {
		return
	}

	id := newGame.UserID
	if existingGame, ok := games[id]; ok {
		if !existingGame.IsFinished() {
			c.IndentedJSON(http.StatusBadRequest, fmt.Sprintf("game exists for user %s", id))
			return
		}
	}
	
	createdGame := engine.NewWithDefaultGuesses(NewWord())
	games[id] = createdGame
	c.IndentedJSON(http.StatusOK, createdGame.ToApiModel(id))
}

func getGame(c *gin.Context) {
	id := c.Param("id")

	var game *engine.Game
	var ok bool
	if game, ok = games[id]; !ok {
		c.IndentedJSON(http.StatusNotFound, fmt.Sprintf("no game exists for user %s", id))
		return
	}
	
	c.IndentedJSON(http.StatusOK, game.ToApiModel(id))
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
	
	if len(newGuess.Guess) != len(game.Target()) {
		errorMsg := fmt.Sprintf(
			"guess must be same length as target word (%d), was %d",
			len(game.Target()),
			len(newGuess.Guess),
		)
		c.IndentedJSON(http.StatusBadRequest, errorMsg)
		return
	}

	if !ValidWord(newGuess.Guess) {
		c.IndentedJSON(http.StatusBadRequest, "guess must be a valid word in the word list")
		return
	}

	game.Guess(newGuess.Guess)
	c.IndentedJSON(http.StatusOK, game.ToApiModel(id))
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "wordle ok")
}
