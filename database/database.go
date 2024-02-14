package database

import (
	"github.com/filipegms5/pokedex-backend/models"

	"gorm.io/gorm"
)

func SetupDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.Pokemon{}, &models.Type{})

	// Check if database is already populated
	if !isDatabasePopulated(db) {
		// Migrate the schema
		createTypes(db)
		createPokemon(db)
	}
}

func isDatabasePopulated(db *gorm.DB) bool {
	var count int64
	db.Model(&models.Pokemon{}).Count(&count)
	return count > 0
}

func createTypes(db *gorm.DB) {
	db.CreateInBatches(TypeList, len(TypeList))
}

func createPokemon(db *gorm.DB) {
	db.CreateInBatches(PokemonList, len(PokemonList))
}
