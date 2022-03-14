package gifts

import (
	"errors"
	"fmt"
	"livegift_back/libraries/response"
	"livegift_back/libraries/storage"
	"livegift_back/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/pop/v5"
)

func CreateGift(c buffalo.Context) error {
	_, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("error trying to connnect to database")
	}

	var gift models.Gift

	c.Request().ParseMultipartForm(10 * 1024 * 1024)
	files, handler, err := c.Request().FormFile("videoFile")
	if err != nil {
		return err
	}

	defer files.Close()

	file := binding.File{files, handler}
	if handler != nil {
		path, err := storage.Upload(c, file)
		if err != nil {
			fmt.Println("======================> Error Upload: ", err)
			return err
		}

		gift.Video = path
	}

	gift.Init(c.Params())
	mapJson := response.Map{
		"gift":    gift,
		"message": "El gift ha sido creado correctamente",
	}

	return response.JSON(c.Response(), c.Request(), http.StatusCreated, mapJson)
}
