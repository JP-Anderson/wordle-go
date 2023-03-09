package rest

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/health", getHealth)	
	return r
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "wordle ok")
}
