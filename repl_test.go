package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		//TestCase: 1
		{
			input: " hELlo World",
			expected: []string{"hello", "world"},
		},
		//TestCase: 2
		{
			input: "",
			expected: []string{},
		},
		//TestCase: 3
		{ 
			input: "  a Quite ComPlicated test   AND    LETS SEEE \n if this  WORKS ",
			expected: []string{"a", "quite", "complicated", "test", "and", "lets", "seee", "if", "this", "works"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("Result Length not matching")
		}

		for i, word := range actual {
			if word != c.expected[i] {
				t.Errorf("Expected and Result is not matching ")
			}
		}
	}
}