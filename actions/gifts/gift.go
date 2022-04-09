package gifts

import (
	"errors"
	"livegift_back/libraries/response"
	"livegift_back/libraries/storage"
	"livegift_back/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
)

func ListGift(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("error trying to connnect to database")
	}

	var gifts models.Gifts
	if err := tx.All(&gifts); err != nil {
		return err
	}

	mapJson := response.Map{
		"gifts": gifts,
	}

	return response.JSON(c.Response(), c.Request(), http.StatusCreated, mapJson)
}

func CreateGift(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("error trying to connnect to database")
	}

	gift := models.Gift{}

	c.Request().ParseMultipartForm(10 * 1024 * 1024)
	files, handler, err := c.Request().FormFile("videoFile")
	if err != nil {
		return err
	}

	defer files.Close()

	file := binding.File{
		File:       files,
		FileHeader: handler,
	}
	if handler != nil {
		secureURL, path, err := storage.Upload(c, file)
		if err != nil {
			return err
		}

		gift.Video = path
		gift.VideoURL = secureURL
	}

	gift.Code = uuid.Must(uuid.NewV4())
	gift.Title = c.Request().FormValue("title")

	if err := tx.Create(&gift); err != nil {
		return err
	}

	mapJson := response.Map{
		"status":  http.StatusCreated,
		"message": "El gift ha sido creado correctamente",
	}

	return response.JSON(c.Response(), c.Request(), http.StatusCreated, mapJson)
}
