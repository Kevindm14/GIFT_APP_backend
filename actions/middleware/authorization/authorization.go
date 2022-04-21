package authorization

import (
	"fmt"
	"livegift_back/libraries/jwt"
	"livegift_back/libraries/response"
	"net/http"
	"os"
	"time"

	"github.com/gobuffalo/buffalo"
)

func Authorizator(next buffalo.Handler) buffalo.Handler {
	signingString := os.Getenv("SIGNING_STRING")

	return func(c buffalo.Context) error {
		c.Logger().Info(">>>> Authorization MW")

		authorization := c.Request().Header.Get("Authorization")
		fmt.Println("======================> ", authorization)

		tokenString, err := TokenFromAuthorization(authorization)
		if err != nil {
			return response.HTTPError(c.Response(), c.Request(), http.StatusUnauthorized, err.Error())
		}

		token, err := jwt.GetFromToken(tokenString, signingString)
		if err != nil {
			return response.HTTPError(c.Response(), c.Request(), http.StatusUnauthorized, err.Error())
		}

		c.Cookies().Set("token", token.ID, time.Duration(time.Now().Add(5*time.Minute).Unix()))
		// response.JSON(c.Response(), c.Request(), http.StatusOK, response.Map{"token": token})

		return next(c)
	}
}
