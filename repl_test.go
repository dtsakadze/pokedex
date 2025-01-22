package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	} {
		{
			input: "  Hello World  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "this is aNOther test inout",
			expected: []string{"this", "is", "another", "test", "input"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word == expectedWord {
				return
			}

			t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
		}
	}
}