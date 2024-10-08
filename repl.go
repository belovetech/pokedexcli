package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/belovetech/pokedexcli.git/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, *pokecache.Cache, *string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show this help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex CLI",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: " Show the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: " Show the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area",
			callback:    commandExplore,
		},
	}
}

func startRepl(cfg *Config, cache *pokecache.Cache) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(" > ")

	for {
		scanner.Scan()

		input := scanner.Text()
		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		availableCommands := getCommands()
		firstArg := ""
		if len(words) > 1 {
			firstArg = words[1]
		}

		command, exists := availableCommands[commandName]

		if !exists {
			fmt.Printf("\nUnknown command: <%v>\n\n", commandName)
			fmt.Print(" > ")
			continue
		}

		// fmt.Println()
		err := command.callback(cfg, cache, &firstArg)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Print("\n > ")
	}
}

func cleanInput(input string) []string {
	output := strings.TrimSpace(input)
	output = strings.ToLower(output)
	return strings.Fields(output)
}
