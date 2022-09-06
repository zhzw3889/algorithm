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
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"234", []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}},
	}
	for _, test := range tests {
		if got := letterCombinations(test.input); !utils.EqualSlice(test.want, got) {
			t.Errorf("letterCombinations(%q) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}

// 全排列问题测试
func TestPermute(t *testing.T) {
	var tests = []struct {
		input []int
		want  [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{2}, [][]int{{2}}},
		{[]int{2, 3}, [][]int{{2, 3}, {3, 2}}},
		{[]int{1, 2, 3}, [][]int{
			{1, 2, 3}, {1, 3, 2},
			{2, 1, 3}, {2, 3, 1},
			{3, 1, 2}, {3, 2, 1}}},
		{[]int{5, 4, 6, 2}, [][]int{
			{5, 4, 6, 2}, {5, 4, 2, 6}, {5, 6, 4, 2}, {5, 6, 2, 4},
			{5, 2, 4, 6}, {5, 2, 6, 4}, {4, 5, 6, 2}, {4, 5, 2, 6},
			{4, 6, 5, 2}, {4, 6, 2, 5}, {4, 2, 5, 6}, {4, 2, 6, 5},
			{6, 5, 4, 2}, {6, 5, 2, 4}, {6, 4, 5, 2}, {6, 4, 2, 5},
			{6, 2, 5, 4}, {6, 2, 4, 5}, {2, 5, 4, 6}, {2, 5, 6, 4},
			{2, 4, 5, 6}, {2, 4, 6, 5}, {2, 6, 5, 4}, {2, 6, 4, 5}}},
	}
	for _, test := range tests {
		if got := permute(test.input); !utils.EqualSlice2Degree(test.want, got) {
			t.Errorf("letterCombinations(%v) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}

// 组合问题测试
func TestCombine(t *testing.T) {
	var tests = []struct {
		input1 int
		input2 int
		want   [][]int
	}{
		{4, 2, [][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
			{2, 4},
			{3, 4},
		}},
		{1, 1, [][]int{{1}}},
	}
	for _, test := range tests {
		if got := combine(test.input1, test.input2); !utils.EqualSlice2Degree(test.want, got) {
			t.Errorf("combine(%v, %v) = %v, wanted: %v", test.input1, test.input2, got, test.want)
		}
	}
}
