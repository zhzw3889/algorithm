package backtracking

// 电话号码的字母组合，难点在于不定长数组的遍历。
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
	// 用于递归的函数
	var findCombination func(digitsIn string, index int, res string)
	findCombination = func(digitsIn string, index int, res string) {
		if index == len(digitsIn) {
			result = append(result, res)
			return
		}
		digit := digitsIn[index]
		letters := letterMp[digit]
		for i := 0; i < len(letters); i++ {
			findCombination(digitsIn, index+1, res+string(letters[i]))
		}
	}
	findCombination(digits, 0, "")
	return result
}

func letterCombinationsLeetcode(digits string) (res []string) {
	if digits == "" {
		return
	}
	mp := make(map[byte]string)
	mp['2'] = "abc"
	mp['3'] = "def"
	mp['4'] = "ghi"
	mp['5'] = "jkl"
	mp['6'] = "mno"
	mp['7'] = "pqrs"
	mp['8'] = "tuv"
	mp['9'] = "wxyz"

	var recur func(str string, count int, ans string)
	recur = func(str string, count int, ans string) {
		if len(str) == count {
			res = append(res, ans)
			return
		}
		s := mp[str[count]]
		for i := 0; i < len(s); i++ {
			recur(str, count+1, ans+string(s[i]))
		}
	}

	recur(digits, 0, "")
	return
}
