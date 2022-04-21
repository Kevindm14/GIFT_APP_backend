package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

type Date time.Time

type Event struct {
	ID     uuid.UUID `json:"id" db:"id"`
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	GiftID uuid.UUID `json:"gift_id" db:"gift_id"`

	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Date        Date   `json:"date" db:"date"`
	Sent        bool   `json:"sent" db:"sent"`

	ParticipantsIDs []uuid.UUID `json:"participants_ids" db:"-"`

	Participants Participants `json:"participants" many_to_many:"event_participants"`
	Gift         Gift         `belongs_to:"gifts" fk_id:"GiftID"`
	User         User         `belongs_to:"users" fk_id:"UserID"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type Events []Event

func (j *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = Date(t)
	return nil
}

func (j Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

// Maybe a Format function for printing your date
func (j Date) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

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
