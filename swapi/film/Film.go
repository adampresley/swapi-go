package film

import "github.com/adampresley/swapi-go/swapi/errors"

type Film struct {
	errors.Error

	Characters   []string `json:"characters"`
	Created      string   `json:"created"`
	Director     string   `json:"director"`
	Edited       string   `json:"edited"`
	EpisodeId    int      `json:"episode_id"`
	OpeningCrawl string   `json:"opening_crawl"`
	Planets      []string `json:"planets"`
	Producer     string   `json:"producer"`
	Species      []string `json:"species"`
	Starships    []string `json:"starships"`
	Title        string   `json:"title"`
	Url          string   `json:"url"`
	Vehicles     []string `json:"vehicles"`
}
