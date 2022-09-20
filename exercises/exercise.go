package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 乘法口诀，右对齐
func printMultipTable() {
	const N = 9
	a := make([][N]string, N)
	for i := 0; i <= N-1; i++ {
		k := 0
		for j := 0; j <= N-1; j++ {
			if i+j >= N-1 {
				jStr := strconv.Itoa(k + 1)
				iStr := strconv.Itoa(i + 1)
				ijStr := strconv.Itoa((k + 1) * (i + 1))
				printStr := fmt.Sprintf("%sX%s=%s", jStr, iStr, ijStr)
				if N <= 9 {
					a[i][j] += fmt.Sprintf("%7s", printStr)
				} else {
					a[i][j] += fmt.Sprintf("%10s", printStr)
				}
				// 每次都要从1开始
				k += 1
			} else {
				if N <= 9 {
					a[i][j] = strings.Repeat(" ", 7)
				} else {
					a[i][j] = strings.Repeat(" ", 10)
				}
			}

		}
	}
	for _, value := range a {
		fmt.Println(value)
	}

}
func main() {
	printMultipTable()
}
