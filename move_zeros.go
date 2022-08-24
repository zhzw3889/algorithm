package main

import "fmt"

func moveZeros(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}
	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	a := []int{0, 2, 3, 0, 4, 5}
	moveZeros(a)
	fmt.Println(a)
}
