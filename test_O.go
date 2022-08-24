package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		n := int(math.Pow10(i))
		begin := time.Now()
		sum := 0
		for i := 0; i < n; i++ {
			sum += i
		}
		end := time.Since(begin)
		fmt.Printf("%-10d\t%v\n", n, end)
	}
}
