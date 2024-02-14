package main

import (
	"os"

	"github.com/filipegms5/pokedex-backend/database"
	"github.com/filipegms5/pokedex-backend/repositories"
	"github.com/filipegms5/pokedex-backend/router"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Instruction os how to run the project on README
func main() {

	// Connect to the database
	db, err := gorm.Open(sqlite.Open("pokedex.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.SetupDatabase(db)

	pokemonRepository := repositories.NewPokemonRepository(db)
	router := router.SetupRouter(pokemonRepository)

	router.Run(os.Getenv("PORT"))
}
