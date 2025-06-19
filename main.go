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

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
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
	}
}

func main() {

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
		cleaned_input := cleanInput(scanner.Text())
		//the first word is the command
		command := cleaned_input[0]
		//get the clicommand and check if it is ok
		cli_command, ok := commands[command]
		//if command is in the list
		if ok {
			cli_command.callback()
		} else {
			fmt.Println("Command not found!")
		}
	}
}

func cleanInput(text string) []string {
	lowercase_text := strings.ToLower(text)
	//The purpose of this function will be to split the users input into "words" based on whitespace.
	// It should also lowercase the input and trim any leading or trailing whitespace. For example:
	split_text := strings.Fields(lowercase_text)
	//return the test
	return split_text
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("commandExit")
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, v := range commands {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	return fmt.Errorf("commandHelp")
}
