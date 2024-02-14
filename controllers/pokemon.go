package controllers

import (
	"net/http"
	"strconv"

	"github.com/filipegms5/pokedex-backend/repositories"

	"github.com/gin-gonic/gin"
)

func GetAllPokemon(c *gin.Context, pokemonRepository *repositories.PokemonRepository) {
	pokemons, err := pokemonRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, pokemons)
}
func GetPokemonById(c *gin.Context, pokemonRepository *repositories.PokemonRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}
	pokemon, err := pokemonRepository.FindOneById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Pokemon not found")
		return
	}

	c.JSON(http.StatusOK, pokemon)
}
