package vehicle

import "github.com/adampresley/swapi-go/swapi/errors"

type Vehicle struct {
	errors.Error

	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	CostInCredits        string   `json:"cost_in_credits"`
	Created              string   `json:"created"`
	Crew                 string   `json:"crew"`
	Edited               string   `json:"edited"`
	Length               string   `json:"length"`
	Manufacturer         string   `json:"manufacturer"`
	MaxAtmospheringSpeed string   `json:"max_atmostphering_speed"`
	Model                string   `json:"model"`
	Name                 string   `json:"name"`
	Passengers           string   `json:"passengers"`
	Films                []string `json:"films"`
	Pilots               []string `json:"pilots"`
	VehicleClass         string   `json:"vehicle_class"`
	Url                  string   `json:"url"`
}
