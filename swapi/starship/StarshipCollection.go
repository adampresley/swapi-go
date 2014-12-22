package starship

import (
	"github.com/adampresley/swapi-go/swapi/collection"
	"github.com/adampresley/swapi-go/swapi/errors"
)

type StarshipCollection struct {
	errors.Error

	collection.SWAPIResultCollection
	Results []Starship `json:"results"`
}
