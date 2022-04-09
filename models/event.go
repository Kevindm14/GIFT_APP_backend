package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

type Event struct {
	ID     uuid.UUID `json:"id" db:"id"`
	GiftID uuid.UUID `json:"gift_id" db:"gift_id"`

	Title       string    `json:"title" db:"title"`
	Description string    `json:"video" db:"video"`
	Date        time.Time `json:"date" db:"date"`
	Sent        bool      `json:"sent" db:"sent"`

	Participants Participants `belongs_to:"users" fk_id:"UserID"`
	Gift         Gift         `belongs_to:"gifts" fk_id:""`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type Events []Event

func (g Event) String() string {
	jg, err := json.Marshal(g)
	if err != nil {
		fmt.Printf("error marshalling json on string nethod: %v\n", err)
	}

	return string(jg)
}

func (g Events) String() string {
	jg, err := json.Marshal(g)
	if err != nil {
		fmt.Printf("error marshalling json on string nethod: %v\n", err)
	}

	return string(jg)
}
