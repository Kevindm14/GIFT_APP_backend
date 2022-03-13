package jwt

import (
	"errors"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofrs/uuid"
)

type Claim struct {
	jwt.StandardClaims
	ID uuid.UUID `json:"id"`
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

	id, ok := iID.(uuid.UUID)
	if !ok {
		return nil, errors.New("invalid user id")
	}

	return &Claim{ID: id}, nil
}
