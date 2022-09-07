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

// 简单选择排序
func SelectSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		// 找到i后面的最小值，记录其下标，与arr[i]交换
		// 比较的次数固定为n(n-1)/2，交换的最大次数为n-1次
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

// 简单插入排序
func InsertionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 保证左侧有序，故不从0开始
	for i := 1; i < len(arr); i++ {
		// 先将arr[i]暂存，待后面放在合适的位置
		tmp := arr[i]
		j := i
		// 保证arr[i]左边的元素是有序的
		// 故只要不满足arr[j-1]>tmp即没有迭代下去的必要
		// 将tmp放在此处(j)即可
		for ; j > 0 && arr[j-1] > tmp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
	return arr
}

// 希尔排序
func ShellSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	N := len(arr)
	var tmp int
	for increment := N / 2; increment > 0; increment /= 2 {
		for i := increment; i < N; i++ {
			tmp = arr[i]
			j := i
			for ; j >= increment; j -= increment {
				if tmp < arr[j-increment] {
					arr[j] = arr[j-increment]
				} else {
					break
				}
			}
			arr[j] = tmp
		}
	}
	return arr
}
