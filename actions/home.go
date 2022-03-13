package actions

import (
	"fmt"
	"livegift_back/actions/middleware/authorization"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	token, err := c.Cookies().Get("token")
	if err := authorization.Verify(err, c); err != nil {
		return errors.Unwrap(err)
	}

	fmt.Println("======>", token)

	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}
