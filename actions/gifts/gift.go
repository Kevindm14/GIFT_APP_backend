package gifts

import (
	"errors"
	"livegift_back/libraries/qr"
	"livegift_back/libraries/response"
	"livegift_back/libraries/storage"
	"livegift_back/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
)

var (
	baseURL = envy.Get("BASE_URL", "http://localhost:3000")
	ENV     = envy.Get("GO_ENV", "development")
)

func GenerateQRCode(c buffalo.Context) error {
	gift := models.Gift{Code: uuid.Must(uuid.NewV4())}
	image, err := qr.CodeForGift(baseURL, gift.Code.String())
	if err != nil {
		return err
	}

	mapJson := response.Map{
		"status":  http.StatusCreated,
		"qrImage": image,
		"giftID":  gift.Code.String(),
	}

	return response.JSON(c.Response(), c.Request(), http.StatusCreated, mapJson)
}

func ListGift(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("error trying to connnect to database")
	}

	gifts := models.Gifts{}
	if err := tx.All(&gifts); err != nil {
		return err
	}

	mapJson := response.Map{
		"gifts": gifts,
	}

	return response.JSON(c.Response(), c.Request(), http.StatusCreated, mapJson)
}

func ListGiftsUsers(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("error trying to connnect to database")
	}

	gifts := models.Gifts{}
	if err := tx.Where("user_id = ?", c.Param("user_id")).All(&gifts); err != nil {
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

	gift.Code = uuid.FromStringOrNil(c.Request().FormValue("giftID"))
	gift.Title = c.Request().FormValue("title")
	gift.Qr = c.Request().FormValue("qrImage")
	gift.UserID = uuid.FromStringOrNil(c.Request().FormValue("userID"))

	if err := tx.Create(&gift); err != nil {
		return err
	}

	mapJson := response.Map{
		"status":  http.StatusCreated,
		"message": "El gift ha sido creado correctamente",
		"gift":    gift,
	}

	return response.JSON(c.Response(), c.Request(), http.StatusCreated, mapJson)
}

func DeleteGift(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("error trying to connnect to database")
	}

	gift := models.Gift{}
	if err := tx.Find(&gift, c.Request().FormValue("gift_id")); err != nil {
		return err
	}

	if err := tx.Destroy(&gift); err != nil {
		return err
	}

	mapJson := response.Map{
		"status":  http.StatusOK,
		"message": "El gift ha sido eliminado correctamente",
	}

	return response.JSON(c.Response(), c.Request(), http.StatusOK, mapJson)
}
