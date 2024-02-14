package router

import (
	"github.com/filipegms5/pokedex-backend/controllers"
	"github.com/filipegms5/pokedex-backend/repositories"
	"github.com/gin-gonic/gin"
)

func SetupRouter(pokemonRepository *repositories.PokemonRepository) *gin.Engine {
	router := gin.Default()

	router.GET("/pokemon", func(c *gin.Context) {
		controllers.GetAllPokemon(c, pokemonRepository)
	})

	router.GET("/pokemon/:id", func(c *gin.Context) {
		controllers.GetPokemonById(c, pokemonRepository)
	})

	return router
}
