package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/InfernoIV/boot.dev-Pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// Update all commands (e.g. help, exit, map) to now accept a pointer to a "config" struct as a parameter.
// This struct will contain the Next and Previous URLs that you'll need to paginate through location areas.
type config struct {
	Next     string
	Current  string
	Previous string
}

var command_list map[string]cliCommand

func init() {
	command_list = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    Command_help,
		},
		"map": {
			name:        "map",
			description: "Get the map",
			callback:    Command_map,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous map",
			callback:    Command_map_back,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    Command_exit,
		},
	}

}

func main() {
	//create configuration
	configuration := config{Current: "https://pokeapi.co/api/v2/location-area", Next: "", Previous: ""}
	//Wait for user input using bufio.NewScanner (this blocks the code and waits for input, once the user types something and presses enter,
	// the code continues and the input is available in the returned bufio.Scanner)
	scanner := bufio.NewScanner(os.Stdin)
	//Start an infinite for loop. This loop will execute once for every command the user types in (we don't want to exit the program after just one command)
	for {
		//Use fmt.Print to print the prompt Pokedex > without a newline character
		fmt.Print("Pokedex >")
		//Use the scanner's .Scan and .Text methods to get the user's input as a string
		scanner.Scan()
		//Clean the user's input string
		cleaned_input := Clean_input(scanner.Text())
		//the first word is the command
		command := cleaned_input[0]
		//get the clicommand and check if it is ok
		cli_command, ok := command_list[command]
		//if command is in the list
		if ok {
			cli_command.callback(&configuration)
		} else {
			fmt.Println("Command not found!")
		}
	}
}

func Clean_input(text string) []string {
	lowercase_text := strings.ToLower(text)
	//The purpose of this function will be to split the users input into "words" based on whitespace.
	// It should also lowercase the input and trim any leading or trailing whitespace. For example:
	split_text := strings.Fields(lowercase_text)
	//return the test
	return split_text
}

func Command_exit(configuration *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func Command_help(configuration *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, v := range command_list {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	return nil
}

func Command_map(configuration *config) error {
	url := configuration.Next
	//if we have no previous url
	if configuration.Next == "" {
		//if we have a current url
		if configuration.Current != "" {
			//use the current url
			url = configuration.Current
		//no urls available
		} else {
			//throw error
			fmt.Printf("No next URL available!\n")
			return nil
		}
	}
	//return the status
	return print_map_data(configuration, url)
}

func Command_map_back(configuration *config) error {
	url := configuration.Previous
	//if we have no previous url
	if configuration.Previous == "" {
		//if we have a current url
		if configuration.Current != "" {
			//use the current url
			url = configuration.Current
		//no urls available
		} else {
			//throw error
			fmt.Printf("No previous URL available!\n")
			return nil
		}
	}
	//return the status
	return print_map_data(configuration, url)
}

func print_map_data(configuration *config, url string) error {
	//get the data
	data, err := pokeapi.Get_map_data(url)
	//if error
	if err != nil {
		//return the error
		return err
	}
	//update the config
	configuration.Current = url
	configuration.Next = data.Next
	configuration.Previous = data.Previous
	//for every result
	for _, v := range data.Results {
		//print
		fmt.Printf("%v\n", v.Name)
	}
	//return
	return nil
}
