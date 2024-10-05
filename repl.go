package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func startRepl() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(" pokedex > ")

	for {
		scanner.Scan()

		input := scanner.Text()
		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		commands := getCommands()

		if command, exists := commands[commandName]; exists {
			err := command.callback()
			if err != nil {
				fmt.Println("Error: ", err)
			}
			fmt.Print(" pokedex > ")
			continue
		} else {
			fmt.Println("Unknown command: ", command)
			continue
		}

	}
}

func cleanInput(input string) []string {
	output := strings.TrimSpace(input)
	output = strings.ToLower(output)
	return strings.Fields(output)
}
