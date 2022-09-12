package sorting

import (
	"fmt"
	"io/ioutil"
	"sorting/utils"
	"strconv"
	"strings"
	"testing"
)

const N int = 100000

// 生成测试[]struct
func generateTests() (tests []struct{ input, want []int }) {
	var a, b, c []int
	// a降序，b升序，c随机
	for i := 1; i <= N; i++ {
		a = append(a, N+1-i)
		b = append(b, i)
	}
	// todo, c从文件中读取，保持统一
	data, err := ioutil.ReadFile("./random_ints.txt")
	if err != nil {
		fmt.Println("Read file err: ", err.Error())
		return
	}
	tmp := strings.Split(string(data), " ")
	for _, v := range tmp {
		num, _ := strconv.Atoi(v)
		c = append(c, num)
	}

	// 测试切片
	tests = []struct {
		input []int
		want  []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{5, 8, 1, 9, 3, 7, 10}, []int{1, 3, 5, 7, 8, 9, 10}},
		{[]int{99, 138, 78, 166, 53}, []int{53, 78, 99, 138, 166}},
		{[]int{99, 138, 78, 166, 53, 1}, []int{1, 53, 78, 99, 138, 166}},
		{a, b}, // 逆序情况
		{b, b}, // 已经排好序
		{c, b}, // 随机序列
	}
	return tests
}

// 冒泡排序01测试
func TestBubbleSort01(t *testing.T) {
	tests := generateTests()
	for _, test := range tests {
		input := make([]int, len(test.input))
		copy(input, test.input)
		if got := BubbleSort01(test.input); !utils.EqualSlice(test.want, got) {
			t.Errorf("BubbleSort01(%v) = %v, wanted: %v", input, got, test.want)
		}
	}
}

// 冒泡排序02测试
func TestBubbleSort02(t *testing.T) {
	tests := generateTests()
	for _, test := range tests {
		input := make([]int, len(test.input))
		copy(input, test.input)
		if got := BubbleSort02(test.input); !utils.EqualSlice(test.want, got) {
			t.Errorf("BubbleSort02(%v) = %v, wanted: %v", input, got, test.want)
		}
	}
}

// 冒泡排序03测试，改进的冒泡
func TestBubbleSort03(t *testing.T) {
	tests := generateTests()
	for _, test := range tests {
		input := make([]int, len(test.input))
		copy(input, test.input)
		if got := BubbleSort03(test.input); !utils.EqualSlice(test.want, got) {
			t.Errorf("BubbleSort03(%v) = %v, wanted: %v", input, got, test.want)
		}
	}
}

// 简单选择排序
func TestSelectSort(t *testing.T) {
	tests := generateTests()
	for _, test := range tests {
		input := make([]int, len(test.input))
		copy(input, test.input)
		if got := SelectSort(test.input); !utils.EqualSlice(test.want, got) {
			t.Errorf("SelectSort(%v) = %v, wanted: %v", input, got, test.want)
		}
	}
}

// 简单插入排序
func TestInsertionSort(t *testing.T) {
	tests := generateTests()
	for _, test := range tests {
		input := make([]int, len(test.input))
		copy(input, test.input)
		if got := InsertionSort(test.input); !utils.EqualSlice(test.want, got) {
			t.Errorf("InsertionSort(%v) = %v, wanted: %v", input, got, test.want)
		}
	}
}

// 希尔排序
func TestShellSort(t *testing.T) {
	tests := generateTests()
	for _, test := range tests {
		input := make([]int, len(test.input))
		copy(input, test.input)
		if got := ShellSort(test.input); !utils.EqualSlice(test.want, got) {
			t.Errorf("ShellSort(%v) = %v, wanted: %v", input, got, test.want)
		}
	}
}

// 堆排序01
func TestHeapSort01(t *testing.T) {
	tests := generateTests()
	for _, test := range tests {
		input := make([]int, len(test.input))
		copy(input, test.input)
		if got := HeapSort01(test.input); !utils.EqualSlice(test.want, got) {
			t.Errorf("HeapSort01(%v) = %v, wanted: %v", input, got, test.want)
		}
	}
}

// 堆排序02
func TestHeapSort02(t *testing.T) {
	tests := generateTests()
	for _, test := range tests {
		input := make([]int, len(test.input))
		copy(input, test.input)
		if got := HeapSort02(test.input); !utils.EqualSlice(test.want, got) {
			t.Errorf("HeapSort02(%v) = %v, wanted: %v", input, got, test.want)
		}
	}
}

// 堆排序03
func TestHeapSort03(t *testing.T) {
	tests := generateTests()
	for _, test := range tests {
		input := make([]int, len(test.input))
		copy(input, test.input)
		if got := HeapSort03(test.input); !utils.EqualSlice(test.want, got) {
			t.Errorf("HeapSort03(%v) = %v, wanted: %v", input, got, test.want)
		}
	}
}
