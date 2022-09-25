package exercises

import (
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
