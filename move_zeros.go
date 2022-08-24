package main

import "fmt"

func moveZeros(nums []int) {
	new_nums := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			new_nums = append(new_nums, nums[i])
		}
	}
	for i := 0; i < len(new_nums); i++ {
		nums[i] = new_nums[i]
	}
	for i := len(new_nums); i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	a := []int{0, 2, 3, 0, 4, 5}
	moveZeros(a)
	fmt.Println(a)
}
