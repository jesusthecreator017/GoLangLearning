package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		commands, exists := getCommands()[commandName]
		if exists {
			err := commands.callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		} else {
			fmt.Println("Unknown command:")
			continue
		}
	}
}

func cleanInput(input string) []string {
	loweredInput := strings.ToLower(input)
	return strings.Fields(loweredInput)
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	Next *string
	Prev *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the map locations",
			callback:    commandMap,
		},
	}
}
