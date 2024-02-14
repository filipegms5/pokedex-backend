package controllers

import (
	"net/http"
	"strconv"

	"github.com/filipegms5/pokedex-backend/repositories"

	"github.com/gin-gonic/gin"
)

func GetAllTypes(c *gin.Context, typeRepository *repositories.TypeRepository) {
	types, err := typeRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, types)
}
func GetTypeByName(c *gin.Context, typeRepository *repositories.TypeRepository) {
	name := c.Param("name")

	typePokemon, err := typeRepository.FindOneByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, "Type not found")
		return
	}

	c.JSON(http.StatusOK, typePokemon)
}
func GetTypeById(c *gin.Context, typeRepository *repositories.TypeRepository) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}
	typePokemon, err := typeRepository.FindOneById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Type not found")
		return
	}

	c.JSON(http.StatusOK, typePokemon)
}
