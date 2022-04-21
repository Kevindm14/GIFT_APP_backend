package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uuid.UUID `json:"id" db:"id"`
	FirstName      string    `json:"first_name" db:"first_name"`
	LastName       string    `json:"last_name" db:"last_name"`
	Email          string    `json:"email" db:"email"`
	Password       string    `json:"password" db:"-"`
	PasswordHash   string    `json:"-" db:"password_hash"`
	PhoneNumber    string    `json:"phone_number" db:"phone_number"`
	PhoneExtension string    `json:"phone_extension" db:"phone_extension"`

	Events Events `has_many:"events"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Users []User

func (u User) String() string {
	ju, err := json.Marshal(u)
	if err != nil {
		fmt.Printf("error marshalling json on string nethod: %v\n", err)
	}

	return string(ju)
}

func (u Users) String() string {
	ju, err := json.Marshal(u)
	if err != nil {
		fmt.Printf("error marshalling json on string nethod: %v\n", err)
	}

	return string(ju)
}

// PasswordMatch compares HashPassword with the password and returns true if they match.
func (u User) PasswordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

	return err == nil
}

func (u *User) BeforeCreate(tx *pop.Connection) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.WithStack(err)
	}

	u.PasswordHash = string(hash)

	return nil
}
