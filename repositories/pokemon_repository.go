package repositories

import (
	"github.com/filipegms5/pokedex-backend/models"
	"gorm.io/gorm"
)

type PokemonRepository struct {
	db *gorm.DB
}

func NewPokemonRepository(db *gorm.DB) *PokemonRepository {
	return &PokemonRepository{db}
}

func (r *PokemonRepository) FindAll() ([]models.Pokemon, error) {
	var pokemonList []models.Pokemon
	err := r.db.Find(&pokemonList).Error
	if err != nil {
		return nil, err
	}
	return pokemonList, err
}
func (r *PokemonRepository) FindOneById(id int) (models.Pokemon, error) {
	var pokemon models.Pokemon
	err := r.db.First(&pokemon, id).Error
	if err != nil {
		return models.Pokemon{}, err
	}
	return pokemon, err
}
