package main

import (
	"fmt"

	"github.com/belovetech/pokedexcli.git/internal/pokecache"
)

func commandExplore(cfg *Config, cache *pokecache.Cache, firstArg *string) error {
	fmt.Printf("Exploring %s...\n", *firstArg)
	pokemen, err := cfg.pokeapiClient.ListPokemonInLocation(*firstArg, cache)
	if err != nil {
		return err
	}

	if len(pokemen) == 0 {
		fmt.Println("No pokemon found in this location")
		return nil
	} else {
		fmt.Println("Found Pokemon:")
	}

	for _, pokeman := range pokemen {
		fmt.Printf(" - %s\n", pokeman)
	}

	return nil
}
