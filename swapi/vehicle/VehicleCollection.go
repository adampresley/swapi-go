package vehicle

import (
	"github.com/adampresley/swapi-go/swapi/collection"
	"github.com/adampresley/swapi-go/swapi/errors"
)

type VehicleCollection struct {
	errors.Error

	collection.SWAPIResultCollection
	Results []Vehicle `json:"results"`
}
