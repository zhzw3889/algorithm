package main

import "testing"

func TestMain(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"{[({})]}", true},
		{"{}", true},
		{"{]", false},
		{"{[}]", false},
		{"{[[[", false},
	}
	for _, test := range tests {
		if got := validParentheses(test.input); got != test.want {
			t.Errorf("validParentheses(%q) = %v", test.input, got)
		}
	}
}
