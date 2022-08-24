package main

import (
	"fmt"
	"time"
)

func binarySearch(sli []int, target int) int {
	length := len(sli)
	l, r := 0, length-1
	for l <= r {
		mid := (l + r) / 2
		if sli[mid] == target {
			return mid
		} else if sli[mid] > target {
			r = sli[mid] - 1
		} else {
			l = sli[mid] + 1
		}
	}
	return -1
}

func main() {
	var a []int
	for i := 0; i < 100000; i++ {
		a = append(a, i)
	}
	targets := []int{
		1, 10, 99, 453, 6666,
	}
	for _, target := range targets {
		start := time.Now()
		index := binarySearch(a, target)
		cost := time.Since(start)
		fmt.Printf("Found %5d in %v: %5d\n", target, cost, index)
	}
}
