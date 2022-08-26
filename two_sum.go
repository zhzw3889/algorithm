package main

import "fmt"

// 有序数组和为定值，返回两者索引
func twoSum(nums []int, target int) [2]int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]+nums[r] == target {
			return [2]int{l, r}
		} else if nums[l]+nums[r] > target {
			r--
		} else {
			l++
		}
	}
	return [2]int{-1, -1}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	// c := []int{5, 4, 3, 2, 1}
	b := twoSum(a, 7)
	fmt.Println(b)
}
