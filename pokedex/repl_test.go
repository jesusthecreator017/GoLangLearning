package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  Hello World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Go is Awesome",
			expected: []string{"go", "is", "awesome"},
		},
		{
			input:    "   Leading and trailing spaces   ",
			expected: []string{"leading", "and", "trailing", "spaces"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("For input '%s', expected %v but got %v", c.input, c.expected, actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("For input '%s', expected word '%s' but got '%s'", c.input, expectedWord, word)
			}
		}
	}
}
