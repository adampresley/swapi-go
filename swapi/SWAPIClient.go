package swapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adampresley/swapi-go/swapi/film"
	"github.com/adampresley/swapi-go/swapi/people"
	"github.com/adampresley/swapi-go/swapi/planet"
	"github.com/adampresley/swapi-go/swapi/root"
	"github.com/adampresley/swapi-go/swapi/species"
	"github.com/adampresley/swapi-go/swapi/starship"
	"github.com/adampresley/swapi-go/swapi/vehicle"
)

var BaseUrl string = "http://swapi.co/api"

type SWAPIClient struct {
}

func NewClient() *SWAPIClient {
	return &SWAPIClient{}
}

func (this *SWAPIClient) buildUrl(endpoint string) string {
	return fmt.Sprintf("%s%s", BaseUrl, endpoint)
}

func (this *SWAPIClient) get(endpoint string) (int, []byte, error) {
	response, err := http.Get(this.buildUrl(endpoint))
	if err != nil {
		return 0, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		return 0, nil, err
	}

	return response.StatusCode, body, nil
}

func (this *SWAPIClient) GetAllFilms() (*film.FilmCollection, int, error) {
	result := &film.FilmCollection{}

	status, body, err := this.get("/films/")
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

func (this *SWAPIClient) GetAllPeople() (*people.PeopleCollection, int, error) {
	result := &people.PeopleCollection{}

	status, body, err := this.get("/people/")
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

func (this *SWAPIClient) GetAllPlanets() (*planet.PlanetCollection, int, error) {
	result := &planet.PlanetCollection{}

	status, body, err := this.get("/planets/")
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

func (this *SWAPIClient) GetAllSpecies() (*species.SpeciesCollection, int, error) {
	result := &species.SpeciesCollection{}

	status, body, err := this.get("/species/")
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

func (this *SWAPIClient) GetAllStarships() (*starship.StarshipCollection, int, error) {
	result := &starship.StarshipCollection{}

	status, body, err := this.get("/starships/")
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

func (this *SWAPIClient) GetAllVehicles() (*vehicle.VehicleCollection, int, error) {
	result := &vehicle.VehicleCollection{}

	status, body, err := this.get("/vehicles/")
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

func (this *SWAPIClient) GetAvailableResources() (*root.Root, int, error) {
	result := &root.Root{}

	status, body, err := this.get("/")
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

func (this *SWAPIClient) GetFilmById(id int) (*film.Film, int, error) {
	result := &film.Film{}

	status, body, err := this.get(fmt.Sprintf("/films/%d", id))
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

func (this *SWAPIClient) GetPersonById(id int) (*people.People, int, error) {
	result := &people.People{}

	status, body, err := this.get(fmt.Sprintf("/people/%d", id))
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

func (this *SWAPIClient) GetPlanetById(id int) (*planet.Planet, int, error) {
	result := &planet.Planet{}

	status, body, err := this.get(fmt.Sprintf("/planets/%d", id))
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

func (this *SWAPIClient) GetSpeciesById(id int) (*species.Species, int, error) {
	result := &species.Species{}

	status, body, err := this.get(fmt.Sprintf("/species/%d", id))
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

func (this *SWAPIClient) GetStarshipById(id int) (*starship.Starship, int, error) {
	result := &starship.Starship{}

	status, body, err := this.get(fmt.Sprintf("/starships/%d", id))
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

func (this *SWAPIClient) GetVehicleById(id int) (*vehicle.Vehicle, int, error) {
	result := &vehicle.Vehicle{}

	status, body, err := this.get(fmt.Sprintf("/vehicles/%d", id))
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}
