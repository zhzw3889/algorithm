package backtracking

import (
	"backtracking/utils"
	"testing"
)

// 电话号码的字母组合
func TestLetterCombinations(t *testing.T) {
	var tests = []struct {
		input string
		want  []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
	}
	for _, test := range tests {
		if got := letterCombinations(test.input); !utils.EqualSlice(test.want, got) {
			t.Errorf("letterCombinations(%q) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}
