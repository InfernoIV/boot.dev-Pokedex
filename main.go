package main

import(
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}


func cleanInput(text string) []string {
	lowercase_text := strings.ToLower(text)
	//The purpose of this function will be to split the users input into "words" based on whitespace. 
	// It should also lowercase the input and trim any leading or trailing whitespace. For example:
	split_text := strings.Fields(lowercase_text)
	//return the test
	return split_text
}