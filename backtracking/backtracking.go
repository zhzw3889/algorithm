package backtracking

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
		// 终止条件，数字序列遍历结束，此时res缓存结束
		if index == len(digitsIn) {
			result = append(result, res)
			return
		}
		// 用index去控制递归层级，进入下一层后首先读取该层的树枝（字符序列）
		digit := digitsIn[index]
		letters := letterMp[digit]
		for i := 0; i < len(letters); i++ {
			findCombination(digitsIn, index+1, res+string(letters[i]))
		}
	}
	// 递归入口
	findCombination(digits, 0, "")
	return result
}
