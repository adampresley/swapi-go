package root

import "github.com/adampresley/swapi-go/swapi/errors"

type Root struct {
	errors.Error

	Films     string `json:"films"`
	People    string `json:"people"`
	Planets   string `json:"planets"`
	Species   string `json:"species"`
	Starships string `json:"starships"`
	Vehicles  string `json:"vehicles"`
}
