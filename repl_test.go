package msf_pokedex

import (
	//"strings"
	"testing"
	//"io"
	//"net/http"
	//"bytes"
	//"encoding/json"
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
			input:    "  more damn noise from this annoying teacher  ",
			expected: []string{"more", "damn", "noise", "from", "this", "annoying", "teacher"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Len(actual):'%d' != Len(expected):'%d'", len(actual), len(c.expected))
			continue // skip the loop that would panic
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test

			//fmt.Printf("Actual:'%s', Expected:'%s'\n", word, expectedWord)

			if word != expectedWord {
				t.Errorf("Actual:'%s' != Expected:'%s'", word, expectedWord)
			}
		}
	}

}
