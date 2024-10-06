package main

import (
	"fmt"
)

func commandMapb(cfg *Config) error {

	location, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationURL)
	if err != nil {
		return err
	}

	fmt.Println("Location names:")
	for _, area := range location.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationURL = location.Next
	cfg.previousLocationURL = location.Previous

	return nil
}
