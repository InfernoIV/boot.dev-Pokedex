package main

import (
	//"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		// add more cases here
	}

	for _, c := range cases {
		
		//convert
		actual := cleanInput(c.input)

		// Check the length of the actual slice against the expected slice
		if(len(actual) != len(c.expected)) {
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			//log the test case input
			t.Logf("Testing: '%v'", c.input)
			t.Errorf("Incorrect lenght! Input: '%v', Expected: '%v'", len(actual), len(c.expected))

		} else {
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				// Check each word in the slice
				if(word != expectedWord) {
					// if they don't match, use t.Errorf to print an error message
					// and fail the test
					//log the test case input
					t.Logf("Testing: '%v'", c.input)
					t.Errorf("Words do not match! '%v', Expected: '%v'", actual, c.expected)
				}
			}
		}
	}
}
