package users

import (
	"errors"
	"livegift_back/libraries/response"
	"livegift_back/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
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
