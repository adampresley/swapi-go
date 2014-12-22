package people

import (
	"github.com/adampresley/swapi-go/swapi/collection"
	"github.com/adampresley/swapi-go/swapi/errors"
)

type PeopleCollection struct {
	errors.Error

	collection.SWAPIResultCollection
	Results []People `json:"results"`
}
