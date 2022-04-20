package events

import (
	"encoding/json"
	"errors"
	"livegift_back/libraries/response"
	"livegift_back/models"
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

func Index(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("error trying to connect database")
	}

	var events models.Events
	if err := tx.All(&events); err != nil {
		log.Println(err)
		return err
	}

	return response.JSON(c.Response(), c.Request(), http.StatusOK, response.Map{"events": events})
}

func Create(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("error trying to connnect to database")
	}

	event := models.Event{}
	if err := json.NewDecoder(c.Request().Body).Decode(&event); err != nil {
		response.HTTPError(c.Response(), c.Request(), http.StatusUnauthorized, err.Error())

		return errors.Unwrap(err)
	}

	if err := tx.Create(&event); err != nil {
		return err
	}

	mapJson := response.Map{
		"status":  http.StatusCreated,
		"message": "El Evento ha sido creado correctamente",
	}

	return response.JSON(c.Response(), c.Request(), http.StatusCreated, mapJson)
}

func DeleteEvent(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("error trying to connnect to database")
	}

	event := models.Event{}
	if err := tx.Find(&event, c.Request().FormValue("event_id")); err != nil {
		return err
	}

	if err := tx.Destroy(&event); err != nil {
		return err
	}

	mapJson := response.Map{
		"status":  http.StatusOK,
		"message": "El Evento ha sido eliminado correctamente",
	}

	return response.JSON(c.Response(), c.Request(), http.StatusOK, mapJson)
}
