package exercises

import (
	"exercise/utils"
	"testing"
)

func TestRotatedDigits(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 1},
		{5, 2},
		{6, 3},
		{7, 3},
		{8, 3},
		{9, 4},
		{10, 4},
		{11, 4},
		{857, 247},
	}
	for _, test := range tests {
		if got := rotatedDigits(test.input); test.want != got {
			t.Errorf("totateDigits(%v) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}

func TestRotatedDigits_zhzw(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		// {1, 0},
		// {2, 1},
		// {3, 1},
		// {4, 1},
		// {5, 2},
		// {6, 3},
		// {7, 3},
		// {8, 3},
		// {9, 4},
		// {10, 4},
		// {11, 4},
		{857, 247},
	}
	for _, test := range tests {
		if got := rotatedDigits_zhzw(test.input); test.want != got {
			t.Errorf("totateDigits(%v) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}

func TestThreeSum(t *testing.T) {
	var tests = []struct {
		input []int
		want  [][]int
	}{
		{nil, [][]int{}},
		{[]int{}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{
			{0, 0, 0},
		}},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		}},
	}
	for _, test := range tests {
		if got := threeSum(test.input); !utils.EqualSlice2Degree(test.want, got) {
			t.Errorf("threeSum(%v) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}
