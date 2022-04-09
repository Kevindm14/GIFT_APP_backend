package events

import (
	"errors"
	"livegift_back/libraries/response"
	"livegift_back/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

func Index(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("error trying to connect database")
	}

	var events models.Events
	if err := tx.All(&events); err != nil {
		return errors.Unwrap(err)
	}

	return response.JSON(c.Response(), c.Request(), http.StatusOK, response.Map{"events": events})
}

func CreateEvent(c buffalo.Context) error {
	return nil
}
