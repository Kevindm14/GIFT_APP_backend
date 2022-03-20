package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type Claim struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func (c *Claim) GetToken(signingString string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(signingString))
}

func GetFromToken(tokenString, signingString string) (*Claim, error) {
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(signingString), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claim")
	}

	iID, ok := claim["id"]
	if !ok {
		return nil, errors.New("user id not found")
	}

	claimObj := Claim{
		fmt.Sprintf("%v", iID),
		jwt.StandardClaims{
			ExpiresAt: &jwt.Time{
				Time: time.Now().AddDate(0, 0, 7),
			},
		},
	}

	return &claimObj, nil
}
