package router

import (
	"github.com/filipegms5/pokedex-backend/controllers"
	"github.com/filipegms5/pokedex-backend/repositories"
	"github.com/gin-gonic/gin"
)

func SetupRouter(pokemonRepository *repositories.PokemonRepository, typeRepository *repositories.TypeRepository) *gin.Engine {
	router := gin.Default()

	router.GET("/pokemon", func(c *gin.Context) {
		controllers.GetAllPokemon(c, pokemonRepository)
	})

	router.GET("/pokemon/:id", func(c *gin.Context) {
		controllers.GetPokemonById(c, pokemonRepository)
	})

	router.GET("/type", func(c *gin.Context) {
		controllers.GetAllTypes(c, typeRepository)
	})

	router.GET("/type/name/:name", func(c *gin.Context) {
		controllers.GetTypeByName(c, typeRepository)
	})
	router.GET("/type/:id", func(c *gin.Context) {
		controllers.GetTypeById(c, typeRepository)
	})

	return router
}
