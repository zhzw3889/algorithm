package sorting

// 冒泡排序01，逆序时，此法比02、03快
func BubbleSort01(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				// arr[i]在i为某个值时一直在变动，一直更新到其在包括自身在内的后续子序列中值为最小
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// 冒泡排序02，普通冒泡
func BubbleSort02(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}

// 冒泡排序03，改进冒泡
func BubbleSort03(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	flag := true
	for i := 0; i < len(arr) && flag; i++ {
		flag = false
		for j := len(arr) - 1; j > i; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				flag = true
			}
		}
	}
	return arr
}
