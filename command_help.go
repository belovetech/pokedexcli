package main

import "fmt"

func commandHelp(cfg *Config) error {
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
