package database

import (
	"github.com/filipegms5/pokedex-backend/models"

	"gorm.io/gorm"
)

func SetupDatabase(db *gorm.DB) {

	// Migrate the schema
	db.AutoMigrate(&models.Pokemon{}, &models.Type{})
	createTypes(db)
	createPokemon(db)

}

func createTypes(db *gorm.DB) {
	db.CreateInBatches(TypeList, len(TypeList))
}

func createPokemon(db *gorm.DB) {
	db.CreateInBatches(PokemonList, len(PokemonList))
}
