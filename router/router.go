package router

import (
	"github.com/filipegms5/pokedex-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/pokemon", func(c *gin.Context) {
		controllers.GetAllPokemon(c)
	})

	router.GET("/pokemon/:id", func(c *gin.Context) {
		controllers.GetPokemonById(c)
	})

	return router
}
