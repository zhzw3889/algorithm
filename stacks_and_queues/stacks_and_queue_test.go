package stacks_and_queues

import "testing"

func TestValidParentheses(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"{[({})]}", true},
		{"{}", true},
		{"[](){}", true},
		{"{]", false},
		{"{[}]", false},
		{"{[[[", false},
		{"{", false},
		{"{}}}", false},
	}
	for _, test := range tests {
		if got := validParentheses(test.input); got != test.want {
			t.Errorf("validParentheses(%q) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}
