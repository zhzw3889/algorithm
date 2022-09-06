package backtracking

import "fmt"

// 电话号码的字母组合，类树问题
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	letterMp := map[byte]string{
		'0': " ",
		'1': "",
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	result := []string{}
	// 递归函数，用递归去缓存result，res为每步递归前的缓存值；go必须先声明，否则无法递归调用
	var findCombination func(digitsIn string, index int, res string)
	findCombination = func(digitsIn string, index int, res string) {
		fmt.Printf("%d : %s\n", index, res)
		// 终止条件，数字序列遍历结束，此时res缓存结束
		if index == len(digitsIn) {
			fmt.Printf("Get %q : take it\n", res)
			result = append(result, res)
			return
		}
		// 用index去控制递归层级，进入下一层后首先读取该层的树枝（字符序列）
		digit := digitsIn[index]
		letters := letterMp[digit]
		for i := 0; i < len(letters); i++ {
			fmt.Printf("digits[%d] = %q, use %q\n", index, digit, letters[i])
			findCombination(digitsIn, index+1, res+string(letters[i]))
		}
	}
	// 递归入口
	findCombination(digits, 0, "")
	return result
}

// Permutations,全排列问题
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	// 不能重复，设标志位
	used := make([]bool, len(nums))
	for i := range nums {
		used[i] = false
	}

	result := [][]int{}
	// 闭包形式时，nums不变量不用作为形参
	var findPermutation func(index int, res []int)
	findPermutation = func(index int, res []int) {
		if index == len(nums) {
			// 这里的copy时必须的！不能修改已经使用的数据，append进result的必须是个全新的slice
			// 上题对应的string，是不可变值
			temp := make([]int, len(nums))
			copy(temp, res)
			result = append(result, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				res = append(res, nums[i])
				used[i] = true
				findPermutation(index+1, res)
				// 要恢复res状态，与上题的不同点
				res = res[:(len(res) - 1)]
				used[i] = false
			}
		}
	}
	findPermutation(0, []int{})
	return result
}

// Combinations，组合问题
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
func combine(n int, k int) [][]int {
	result := [][]int{}
	if n <= 0 || k <= 0 || k > n {
		return result
	}
	// 循环开始条件为start
	var generateCombinations func(start int, c []int)
	generateCombinations = func(start int, c []int) {
		// 还有k-len(c)个空位
		if n+1-(k-len(c)) < 0 {
			return
		}

		if len(c) == k {
			temp := make([]int, k)
			copy(temp, c)
			result = append(result, temp)
		}
		for i := start; i <= n; i++ {
			c = append(c, i)
			generateCombinations(i+1, c)
			c = c[:len(c)-1]
		}
	}
	generateCombinations(1, []int{})
	return result
}
