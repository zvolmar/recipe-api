package handlers

import (
	"nomcom-api/cmd/models"
	"nomcom-api/cmd/repos"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateRecipe(c echo.Context) error {
	recipe := models.Recipe{}
	c.Bind(&recipe)
	newRecipe, err := repos.CreateRecipe(recipe)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newRecipe)
}

func UpdateRecipe(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
		err.Error())
	}

	recipe := models.Recipe{}
	c.Bind(&recipe)
	updatedRecipe, err := repos.UpdateRecipe(recipe, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedRecipe)
}

func GetRecipe(c echo.Context) error {
	id := c.Param("id")
	
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
		err.Error())
	}

	recipe := models.Recipe{}
	c.Bind(&recipe)
	desiredRecipe, err := repos.GetRecipe(recipe, idInt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, desiredRecipe)
}

// func GetAllRecipesByUser(c echo.Context) error {
// 	id := c.Param("user_id")

// 	idInt, err := strconv.Atoi(id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError,
// 		err.Error())
// 	}

// 	recipes := []models.Recipe{}
// 	c.Bind(&recipes)
// 	recipeList, err := repos.GetAllRecipesByUser(idInt)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, recipeList)
// }