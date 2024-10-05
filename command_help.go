package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex CLI")
	fmt.Printf("Usage: \n\n")

	fmt.Println("help - Show this help message")
	fmt.Println("exit - Exit the Pokedex CLI")

	return nil
}
