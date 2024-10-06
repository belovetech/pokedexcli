package main

import "github.com/belovetech/pokedexcli.git/internal/pokeapi"

type Config struct {
	pokeapiClient       *pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
}

func main() {
	cfg := Config{
		pokeapiClient: &pokeapi.Client{},
	}

	startRepl(&cfg)
}
