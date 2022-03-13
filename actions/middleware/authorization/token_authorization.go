package authorization

import (
	"errors"
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

	return l[1], nil
}
