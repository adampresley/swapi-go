package planet

import (
	"github.com/adampresley/swapi-go/swapi/collection"
	"github.com/adampresley/swapi-go/swapi/errors"
)

type PlanetCollection struct {
	errors.Error

	collection.SWAPIResultCollection
	Results []Planet `json:"results"`
}
