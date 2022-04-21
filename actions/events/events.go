package events

import (
	"encoding/json"
	"errors"
	"livegift_back/actions/middleware/authorization"
	"livegift_back/libraries/response"
	"livegift_back/models"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
)

func Index(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("error trying to connect database")
	}

	authorizationHeader := c.Request().Header.Get("Authorization")
	token, err := authorization.TokenFromAuthorization(authorizationHeader)
	if err != nil {
		return err
	}

	claims := jwt.MapClaims{}
	_, err2 := jwt.ParseWithClaims(token, claims, func(tokenMap *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNING_STRING")), nil
	})

	if err2 != nil {
		return err
	}

	events := models.Events{}
	if err := tx.Where("user_id = ?", claims["id"]).EagerPreload("Gift", "User", "User.Events").All(&events); err != nil {
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
	event.UserID = uuid.FromStringOrNil(c.Request().FormValue("user_id"))
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
