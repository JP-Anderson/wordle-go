package main

import (
	"wordle/rest"	
)

func main() {
	router := rest.Router()
	router.Run("localhost:8080")
}

