package species

import (
	"github.com/adampresley/swapi-go/swapi/collection"
	"github.com/adampresley/swapi-go/swapi/errors"
)

type SpeciesCollection struct {
	errors.Error

	collection.SWAPIResultCollection
	Results []Species `json:"results"`
}
