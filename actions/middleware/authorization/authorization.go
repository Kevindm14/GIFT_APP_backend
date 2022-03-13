package authorization

import (
	"livegift_back/libraries/jwt"
	"livegift_back/libraries/response"
	"net/http"
	"os"

	"github.com/gobuffalo/buffalo"
)

func Authorizator(next buffalo.Handler) buffalo.Handler {
	signingString := os.Getenv("SIGNING_STRING")

	return func(c buffalo.Context) error {
		c.Logger().Info(">>>> Authorization MW")

		authorization := c.Request().Header.Get("Authorization")
		tokenString, err := tokenFromAuthorization(authorization)
		if err != nil {
			return response.HTTPError(c.Response(), c.Request(), http.StatusUnauthorized, err.Error())
		}

		token, err := jwt.GetFromToken(tokenString, signingString)
		if err != nil {
			return response.HTTPError(c.Response(), c.Request(), http.StatusUnauthorized, err.Error())
		}

		c.Set("token", token)
		// response.JSON(c.Response(), c.Request(), http.StatusOK, response.Map{"token": token})

		return next(c)
	}
}
