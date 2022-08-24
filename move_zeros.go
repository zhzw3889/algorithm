package main

import "fmt"

func moveZeros(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// 防止自身与自身交换
			if i != k {
				nums[k], nums[i] = nums[i], nums[k]
			}
			// 元素为0时k值暂停自增，其余时候与i值同步增加
			k++
		}
	}
}

func main() {
	a := []int{2, 0, 4, 5}
	moveZeros(a)
	fmt.Println(a)
}
