package main

import (
	"time"

	"github.com/belovetech/pokedexcli.git/internal/pokeapi"
	"github.com/belovetech/pokedexcli.git/internal/pokecache"
)

type Config struct {
	pokeapiClient       *pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
}

func main() {
	cfg := Config{
		pokeapiClient: &pokeapi.Client{},
	}

	interval := 5 * time.Minute // 5 minutes
	cache := pokecache.NewCache(interval)

	startRepl(&cfg, cache)

}
