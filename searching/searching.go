package searching

// 二分查找，针对静态查找表，时间复杂度为O(logn)
// 注意边界问题
func BinarySearch(nums []int, target int) int {
	N := len(nums)
	low, high := 0, N-1
	var mid int
	// 当low=high时只有一个元素，mid=high=low
	// 下一步如果还不满足target==nums[mid]，则会出现low>high的情况，跳出循环
	for low <= high {
		// mid = (low + high) / 2
		// 避免整型溢出
		mid = low + (high-low)/2
		if target < nums[mid] {
			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// 插值查找，改进mid取值
func InterpolationSearch(nums []int, target int) int {
	N := len(nums)
	low, high := 0, N-1
	var mid int
	for low <= high {
		// 只有此处改进，mid可能出界
		// 1/2变成(target-nums[low])/(nums[high]-nums[low])
		mid = int(float64(low) + float64(target-nums[low])/float64(nums[high]-nums[low])*float64(high-low))
		if mid < 0 || mid > N-1 {
			break
		}
		if target < nums[mid] {
			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
