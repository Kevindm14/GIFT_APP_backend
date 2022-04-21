package users

import (
	"encoding/json"
	"errors"
	"livegift_back/libraries/response"
	"livegift_back/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Index(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("error trying to connect database")
	}

	var users models.Users
	if err := tx.All(&users); err != nil {
		return errors.Unwrap(err)
	}

	return response.JSON(c.Response(), c.Request(), http.StatusOK, response.Map{"users": users})
}

func Update(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("error trying to connect database")
	}

	user := models.User{}
	if err := tx.Find(&user, c.Param("user_id")); err != nil {
		return response.HTTPError(c.Response(), c.Request(), http.StatusUnauthorized, err.Error())
	}

	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return response.HTTPError(c.Response(), c.Request(), http.StatusUnauthorized, err.Error())
	}

	if err := tx.Update(&user); err != nil {
		return response.HTTPError(c.Response(), c.Request(), http.StatusUnauthorized, err.Error())
	}

	mapResponse := response.Map{
		"status":  http.StatusOK,
		"message": "El usuario ha sido eliminado correctamente",
	}

	return response.JSON(c.Response(), c.Request(), http.StatusOK, mapResponse)
}

func Delete(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("error trying to connect database")
	}

	user := models.User{}
	if err := tx.Find(&user, c.Param("user_id")); err != nil {
		return response.HTTPError(c.Response(), c.Request(), http.StatusUnauthorized, err.Error())
	}

	if err := tx.Destroy(&user); err != nil {
		return response.HTTPError(c.Response(), c.Request(), http.StatusUnauthorized, err.Error())
	}

	mapResponse := response.Map{
		"status":  http.StatusOK,
		"message": "El usuario ha sido eliminado correctamente",
	}

	return response.JSON(c.Response(), c.Request(), http.StatusOK, mapResponse)
}
