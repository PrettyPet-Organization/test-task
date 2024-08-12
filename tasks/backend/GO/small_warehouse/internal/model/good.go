package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

const imageProvider = "https://api.waifu.pics/%s/%s"

type Model interface {
	GetKey() string
}

type Good struct {
	Id       uuid.UUID `json:"id,omitempty"`
	Name     string    `json:"name"`
	Desc     string    `json:"description"`
	ImageURL string    `json:"image_url,omitempty"`
}

func NewGood(typ, category string) *Good {
	return &Good{
		Id:       uuid.New(),
		Name:     gofakeit.ProductName(),
		Desc:     gofakeit.ProductDescription(),
		ImageURL: getImageURL(typ, category),
	}
}

func (g Good) GetKey() string {
	return g.Id.String()
}

func getImageURL(typ, category string) string {
	if typ == "" {
		typ = "nsfw"
	}
	if category == "" {
		category = "waifu"
	}

	url := fmt.Sprintf(imageProvider, typ, category)

	r, err := http.Get(url)
	if err != nil {
		return ""
	}

	type ImageProviderResponce struct {
		Url string `json:"url"`
	}

	bod, err := io.ReadAll(r.Body)
	if err != nil {
		return ""
	}

	var m ImageProviderResponce
	err = json.Unmarshal(bod, &m)
	if err != nil {
		return ""
	}

	return m.Url
}
