package controllers

import (
	"net/http"
	"strconv"

	"github.com/filipegms5/pokedex-backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllPokemon(c *gin.Context) {
	pokemons := models.FetchAllPokemon()
	c.JSON(http.StatusOK, pokemons)
}
func GetPokemonById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}
	pokemon := models.FetchPokemonById(id)
	if pokemon.Id == 0 {
		c.JSON(http.StatusNotFound, "Pokemon not found")
		return
	}
	c.JSON(http.StatusOK, pokemon)
}
