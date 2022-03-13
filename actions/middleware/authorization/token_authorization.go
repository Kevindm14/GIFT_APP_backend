package authorization

import (
	"errors"
	"fmt"
	"strings"
)

func tokenFromAuthorization(authorization string) (string, error) {
	if authorization == "" {
		return "", errors.New("authorization is required")
	}

	if !strings.HasPrefix(authorization, "Bearer") {
		return "", errors.New("invalid authorization format")
	}

	l := strings.Split(authorization, " ")
	if len(l) != 2 {
		return "", errors.New("invalid authorization format")
	}

	fmt.Println("============>TOKEN: ", l[1])

	return l[1], nil
}
