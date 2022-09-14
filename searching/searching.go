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
		mid = (low + high) / 2
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
