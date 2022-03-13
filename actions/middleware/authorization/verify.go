package authorization

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func Verify(err error, ctx buffalo.Context) error {
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			ctx.Response().WriteHeader(http.StatusUnauthorized)

			return err
		}
		// For any other type of error, return a bad request status
		ctx.Response().WriteHeader(http.StatusUnauthorized)

		return err
	}

	return nil
}
