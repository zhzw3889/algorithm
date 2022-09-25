package exercises

// todo
func rotatedDigits_zhzw(n int) int {
	// false为翻转不同，true为翻转相同
	digits := map[int]bool{
		0: true,
		1: true,
		2: false,
		5: false,
		6: false,
		8: true,
		9: false,
	}

	result := 0
	var isGoodNum func(int, bool) bool
	isGoodNum = func(i int, haveDiff bool) bool {
		digit := i % 10
		// 3,4,7
		if _, ok := digits[digit]; !ok {
			return false
		}
		// 2,5,6,9
		if diffDigit := digits[digit]; !diffDigit {
			haveDiff = true
		}
		if i < 10 {
			return haveDiff
		}

		isGoodNum(i/10, haveDiff)
		return false
	}

	// 所有数字都在digits中，至少要含有一个false的数
	for i := 1; i <= n; i++ {
		if isGoodNum(i, false) {
			result++
		}
	}
	return result
}

// 旋转数字，双百答案，未用递归
func rotatedDigits(n int) int {
	isRotatedDigits := func(x int) bool {
		flag := false
		for ; x > 0; x /= 10 {
			switch x % 10 {
			case 3, 4, 7:
				return false
			case 2, 5, 6, 9:
				flag = true
			}
		}
		return flag
	}
	res := 0
	for i := 1; i <= n; i++ {
		if isRotatedDigits(i) {
			res++
		}
	}
	return res
}
