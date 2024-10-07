package main

import (
	"fmt"

	"github.com/belovetech/pokedexcli.git/internal/pokecache"
)

func commandHelp(cfg *Config, cache *pokecache.Cache) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex CLI")
	fmt.Println("Usage: ")
	fmt.Println()

	for _, command := range getCommands() {
		fmt.Printf("  %s: %s\n", command.name, command.description)
	}

	fmt.Println()
	return nil
}
