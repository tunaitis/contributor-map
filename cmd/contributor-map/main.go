package main

import (
	"errors"
	"github.com/tunaitis/contributor-map/internal/github"
	"github.com/tunaitis/contributor-map/internal/location"
	"github.com/tunaitis/contributor-map/internal/render"
	"log"
	"os"
	"strings"
)

type config struct {
	useCache    bool
	repository  string
	output      string
	accessToken string
	palette     []string
}

func getConfig() (*config, error) {
	c := config{}

	c.useCache = false
	if os.Getenv("INPUT_CACHE") != "" {
		c.useCache = true
	}

	c.palette = []string{"#D9E887", "#B3D778", "#8CC569", "#6AB356", "#47A042", "#328533", "#1D6A23"}
	if os.Getenv("INPUT_PALETTE") != "" {
		c.palette = strings.Split(os.Getenv("INPUT_PALETTE"), ",")
	}

	c.repository = os.Getenv("INPUT_REPOSITORY")
	if c.repository == "" {
		return nil, errors.New("repository name is not provided")
	}

	c.output = os.Getenv("INPUT_OUTPUT")
	if c.output == "" {
		c.output = "render.svg"
	}

	c.accessToken = os.Getenv("INPUT_TOKEN")
	if c.accessToken == "" {
		return nil, errors.New("GitHub access token is not provided")
	}

	return &c, nil
}

func main() {
	log.SetPrefix("contributor-map: ")
	cfg, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("starting with the repository: %s", cfg.repository)

	contributors, err := github.GetContributors(cfg.accessToken, cfg.repository, 1, cfg.useCache)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("found %d contributors", len(contributors))

	contributors, err = github.GetLocations(cfg.accessToken, contributors, cfg.useCache)
	if err != nil {
		log.Fatal(err)
	}

	db, err := location.NewDb()
	if err != nil {
		log.Fatal(err)
	}

	countries := map[string]int{}
	hasLocation := 0
	for _, c := range contributors {
		country := db.Search(c.Location)
		if country == "" {
			continue
		}

		hasLocation += 1

		if val, ok := countries[country]; ok {
			countries[country] = val + c.Contributions
		} else {
			countries[country] = c.Contributions
		}
	}

	log.Printf("found %d locations", hasLocation)

	log.Println("generating map")
	svg, err := render.Map(countries, cfg.palette)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("writing map to file: %s", cfg.output)
	err = os.WriteFile(cfg.output, svg, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
