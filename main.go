package main

import (
	"os"

	"github.com/filipegms5/pokedex-backend/router"
)

// Instruction os how to run the project on README
func main() {
	router := router.SetupRouter()

	router.Run(os.Getenv("PORT"))
}
