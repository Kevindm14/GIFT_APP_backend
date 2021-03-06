package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/binding"
	"github.com/gofrs/uuid"
)

type Gift struct {
	ID     uuid.UUID `json:"id" db:"id"`
	Code   uuid.UUID `json:"code" db:"code"`
	UserID uuid.UUID `json:"user_id" db:"user_id"`

	Title    string `json:"title" db:"title"`
	Video    string `json:"video" db:"video"`
	VideoURL string `json:"videoURL" db:"video_url"`
	Qr       string `json:"qr" db:"qr"`

	VideoFile binding.File `json:"-" db:"-"`
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

func (g *Gift) Init(params buffalo.ParamValues) {
	g.Code = uuid.FromStringOrNil(params.Get("code"))
	g.Title = params.Get("title")
}
