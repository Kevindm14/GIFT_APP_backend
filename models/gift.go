package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gobuffalo/buffalo/binding"
	"github.com/gofrs/uuid"
)

type Gift struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Code uuid.UUID `json:"code" db:"code"`

	Title     string `json:"title" db:"title"`
	Video     string `json:"video" db:"video"`
	Reference int    `json:"reference" db:"reference" rw:"r"`
	VideoURL  string `json:"videoURL" db:"video_url"`
	Qr        string `json:"QR" db:"-"`

	VideoFile binding.File `json:"videoFile" db:"-"`
	ImageFile string       `json:"imageFile" db:"-"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type Gifts []Gift

func (g Gift) String() string {
	jg, err := json.Marshal(g)
	if err != nil {
		fmt.Printf("error marshalling json on string nethod: %v\n", err)
	}

	return string(jg)
}

func (g Gifts) String() string {
	jg, err := json.Marshal(g)
	if err != nil {
		fmt.Printf("error marshalling json on string nethod: %v\n", err)
	}

	return string(jg)
}
