package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/health", getHealth)
	router.Run("localhost:8080")
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "ok")
}
