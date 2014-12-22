package planet

import "github.com/adampresley/swapi-go/swapi/errors"

type Planet struct {
	errors.Error

	Climate        string   `json:"climate"`
	Created        string   `json:"created"`
	Diameter       string   `json:"diameter"`
	Edited         string   `json:"edited"`
	Films          []string `json:"films"`
	Gravity        string   `json:"gravity"`
	Name           string   `json:"name"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Population     string   `json:"population"`
	Residents      []string `json:"residents"`
	RotationPeriod string   `json:"rotation_period"`
	SurfaceWater   string   `json:"surface_water"`
	Terrain        string   `json:"terrain"`
	Url            string   `json:"url"`
}
