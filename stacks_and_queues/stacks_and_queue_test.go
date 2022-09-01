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

// 逆波兰表达式测试
func TestEvalRPN(t *testing.T) {
	var tests = []struct {
		input []string
		want  int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for _, test := range tests {
		if got := evalRPN(test.input); got != test.want {
			t.Errorf("evalRPN(%q) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}
