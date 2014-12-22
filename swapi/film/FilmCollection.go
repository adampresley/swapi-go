package film

import (
	"github.com/adampresley/swapi-go/swapi/collection"
	"github.com/adampresley/swapi-go/swapi/errors"
)

type FilmCollection struct {
	errors.Error

	collection.SWAPIResultCollection
	Results []Film `json:"results"`
}
