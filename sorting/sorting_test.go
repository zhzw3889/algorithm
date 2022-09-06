package sorting

import (
	"sorting/utils"
	"testing"
)

const N int = 100000

// 生成测试[]struct
func generateTests() (tests []struct{ input, want []int }) {
	var a, b []int
	// a降序，b升序
	for i := 1; i <= N; i++ {
		a = append(a, N+1-i)
		b = append(b, i)
	}
	tests = []struct {
		input []int
		want  []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{9, 8, 1, 5, 3}, []int{1, 3, 5, 8, 9}},
		{[]int{99, 138, 78, 166, 53}, []int{53, 78, 99, 138, 166}},
		{a, b},
	}
	return tests
}

// 冒泡排序01测试, 3.24s
func TestBubbleSort01(t *testing.T) {
	tests := generateTests()
	for _, test := range tests {
		if got := BubbleSort01(test.input); !utils.EqualSlice(test.want, got) {
			t.Errorf("BubbleSort01(%v) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}

// 冒泡排序02测试, 8.90s
func TestBubbleSort02(t *testing.T) {
	tests := generateTests()
	for _, test := range tests {
		if got := BubbleSort02(test.input); !utils.EqualSlice(test.want, got) {
			t.Errorf("BubbleSort02(%v) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}

// 冒泡排序03测试，改进的冒泡, 8.52s
func TestBubbleSort03(t *testing.T) {
	tests := generateTests()
	for _, test := range tests {
		if got := BubbleSort03(test.input); !utils.EqualSlice(test.want, got) {
			t.Errorf("BubbleSort03(%v) = %v, wanted: %v", test.input, got, test.want)
		}
	}
}
