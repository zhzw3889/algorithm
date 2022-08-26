package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, -1
	sum := 0
	res := len(nums) + 1
	for l < len(nums) {
		if r+1 < len(nums) && sum < target {
			r++
			sum += nums[r]
		} else {
			sum -= nums[l]
			l++
		}

		if sum >= target && res > r-l+1 {
			res = r - l + 1
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(minSubArrayLen(7, a))
	fmt.Println(minSubArrayLen(15, a))
}
