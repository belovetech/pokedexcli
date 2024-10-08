package main

import (
	"os"

	"github.com/belovetech/pokedexcli.git/internal/pokecache"
)

func commandExit(cfg *Config, cache *pokecache.Cache, firstArg *string) error {
	os.Exit(0)
	return nil
}
