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

// 堆排序01 (大话数据结构)
func HeapSort01(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	N := len(arr)
	// 将完全二叉树调整成大顶堆
	HeapAdjust := func(s, m int) {
		temp := arr[s]
		for j := 2 * s; j <= m; j *= 2 {
			if j < m && arr[j] < arr[j+1] {
				j++
			}
			if temp >= arr[j] {
				break
			}
			arr[s] = arr[j]
			s = j
		}
		arr[s] = temp
	}

	for i := N / 2; i >= 0; i-- {
		HeapAdjust(i, N-1)
	}
	for i := N - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		HeapAdjust(0, i-1)
	}

	return arr
}

// 堆排序02，（DSAAC）
func HeapSort02(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	N := len(arr)
	// 将完全二叉树调整成大顶堆，n为堆长度
	heapify := func(j, n int) {
		var tmp, child int
		for tmp = arr[j]; 2*j+1 < n; j = child {
			child = 2*j + 1
			// child != n-1为边界条件，保证只在树内操作
			if child != n-1 && arr[child+1] > arr[child] {
				child++
			}
			if tmp < arr[child] {
				arr[j] = arr[child]
			} else {
				// 默认结尾是已经堆化的
				// 与构建时倒序迭代是相关的
				break
			}
		}
		arr[j] = tmp
	}
	// 从最结尾的父节点开始
	for i := N / 2; i >= 0; i-- {
		// 与break相关
		heapify(i, N)
	}

	// 将根节点与尾节点交换，再重新调整为大顶堆(maxheap)
	for i := N - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(0, i)
	}

	return arr
}

// 堆排序03，（递归heapify）
func HeapSort03(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	N := len(arr)
	// 用递归实现堆化，理解上更简单一些
	var heapify func(j, n int)
	heapify = func(j, n int) {
		leftChild := 2*j + 1
		rightChild := 2*j + 2
		max := j
		if leftChild <= n-1 && arr[leftChild] > arr[max] {
			max = leftChild
		}
		if rightChild <= n-1 && arr[rightChild] > arr[max] {
			max = rightChild
		}
		if max != j {
			arr[j], arr[max] = arr[max], arr[j]
			heapify(max, n)
		} else {
			// 与前二者break作用同，创建时可以节省一半时间
			return
		}
	}

	for i := N / 2; i >= 0; i-- {
		heapify(i, N)
	}

	for i := N - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(0, i)
	}

	return arr
}

// 归并排序，递归实现
func MergeSort01(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	N := len(arr)

	merge := func(tmpArray []int, lpos, rpos, rightEnd int) {
		var leftEnd, numElements, tmpPos int
		leftEnd = rpos - 1
		tmpPos = lpos
		numElements = rightEnd - lpos + 1

		for lpos <= leftEnd && rpos <= rightEnd {
			// main loop
			if arr[lpos] <= arr[rpos] {
				tmpArray[tmpPos] = arr[lpos]
				lpos++
				tmpPos++
			} else {
				tmpArray[tmpPos] = arr[rpos]
				rpos++
				tmpPos++
			}
		}
		// copy rest of first half
		for lpos <= leftEnd {
			tmpArray[tmpPos] = arr[lpos]
			lpos++
			tmpPos++
		}
		// copy rest of second half
		for rpos <= rightEnd {
			tmpArray[tmpPos] = arr[rpos]
			rpos++
			tmpPos++
		}
		// copy tmpArray back
		for i := 0; i < numElements; i++ {
			arr[rightEnd] = tmpArray[rightEnd]
			rightEnd--
		}
	}

	var mSort func(tmpArray []int, left, right int)
	mSort = func(tmpArray []int, left, right int) {
		var center int
		if left < right {
			center = (left + right) / 2
			mSort(tmpArray, left, center)
			mSort(tmpArray, center+1, right)
			merge(tmpArray, left, center+1, right)
		}
	}

	tmpArray := make([]int, N)
	mSort(tmpArray, 0, N-1)
	return arr
}

// 快速排序01，第一个元素作为pivot
func QuickSort01(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	N := len(arr)

	partition := func(low, high int) int {
		// 选第一个数作为枢轴值
		pivotval := arr[low]
		for low < high {
			for low < high && arr[high] >= pivotval {
				high--
			}
			arr[low], arr[high] = arr[high], arr[low]
			for low < high && arr[low] <= pivotval {
				low++
			}
			arr[low], arr[high] = arr[high], arr[low]
		}
		return low
	}

	var qSort func(low, high int)
	qSort = func(low, high int) {
		var pivot int
		if low < high {
			pivot = partition(low, high)
			qSort(low, pivot-1)
			qSort(pivot+1, high)
		}
	}
	qSort(0, N-1)
	return arr
}
