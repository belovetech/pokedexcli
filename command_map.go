package main

import (
	"fmt"

	"github.com/belovetech/pokedexcli.git/internal/pokecache"
)

func commandMap(cfg *Config, cache *pokecache.Cache) error {

	location, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL, cache)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range location.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationURL = location.Next
	cfg.previousLocationURL = location.Previous

	return nil
}
