package main

import (
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
			input:    "HelloWorld",
			expected: []string{"helloworld"},
		},
		{
			input:    "Hello  World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello world",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Wrong number of words. Expected %v, but actual is %v.", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Wrong word. Expected %v, but actual` is %v.", expectedWord, word)
			}
		}
	}
}
