package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

type Participant struct {
	ID uuid.UUID `json:"id" db:"id"`

	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type Participants []Participant

func (g Participant) String() string {
	jg, err := json.Marshal(g)
	if err != nil {
		fmt.Printf("error marshalling json on string nethod: %v\n", err)
	}

	return string(jg)
}

func (g Participants) String() string {
	jg, err := json.Marshal(g)
	if err != nil {
		fmt.Printf("error marshalling json on string nethod: %v\n", err)
	}

	return string(jg)
}
