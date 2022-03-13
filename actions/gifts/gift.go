package gifts

import (
	"encoding/json"
	"errors"
	"livegift_back/libraries/response"
	"livegift_back/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

func CreateGift(c buffalo.Context) error {
	_, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("error trying to connnect to database")
	}

	var gift models.Gift
	if err := json.NewDecoder(c.Request().Body).Decode(&gift); err != nil {
		response.HTTPError(c.Response(), c.Request(), http.StatusBadRequest, err.Error())

		return errors.Unwrap(err)
	}

	mapJson := response.Map{
		"gift":    gift,
		"message": "El gift ha sido creado correctamente",
	}

	return response.JSON(c.Response(), c.Request(), http.StatusCreated, mapJson)
}
