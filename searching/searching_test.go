package searching

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		input01 []int
		input02 int
		want    int
	}{
		{[]int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}, 24, 3},   // 靠前
		{[]int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}, 62, 7},   // 靠后
		{[]int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}, 0, 0},    // 首
		{[]int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}, 99, 10},  // 尾
		{[]int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}, 999, -1}, // 无
	}
	for _, test := range tests {
		if got := BinarySearch(test.input01, test.input02); test.want != got {
			t.Errorf("BinarySearch(%v, %v) = %v, wanted: %v", test.input01, test.input02, got, test.want)
		}
	}
}
