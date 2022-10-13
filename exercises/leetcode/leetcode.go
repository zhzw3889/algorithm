package exercises

import (
	"exercise/utils"
	"math"
	"sort"
)

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

// 两数相加
func addTwoNumbers_zhzw(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	l := &utils.ListNode{}
	// 进位符
	flag := 0
	for {
		l.Val = l1.Val + l2.Val
		if flag == 1 {
			l.Val++
		}
		if l.Val >= 10 {
			l.Val /= l.Val
			flag = 1
		} else {
			flag = 0
		}
		l.Next = &utils.ListNode{}
		if l.Next == nil {
			break
		}
		l = l.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	return l
}

// 两数相加
func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	var head *utils.ListNode
	var tail *utils.ListNode
	// 余数
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		// sum%10是节点的当前值，如果是10,取余后当前节点值为0，sum/10是求十位的那个数
		sum, carry = sum%10, sum/10
		// 此时申请一个新的链表存储两个链表的和
		if head == nil {
			// 申请新的链表
			head = &utils.ListNode{Val: sum}
			// 这一步是为了保持头结点不变的情况下指针可以右移，所以说tail相当于临时节点，理解成尾节点也可以，因
			// 为此时新链表中只有一个节点，所以头结点和尾结点都指向同一个元素。
			tail = head
		} else {
			// 第二个节点后开始逐渐往尾结点增加元素
			tail.Next = &utils.ListNode{Val: sum}
			tail = tail.Next
		}
	}
	// 把最后一位的余数加到链表最后。
	if carry > 0 {
		tail.Next = &utils.ListNode{Val: carry}
	}
	return head
}

// 整数翻转
// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
// 如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。
func reverse_zhzw(x int) int {
	var y int
	if x == 0 {
		return x
	} else if x > 0 {
		for x > 0 {
			y = 10*y + x%10
			x /= 10
		}
	} else {
		x = -x
		for x > 0 {
			y = 10*y + x%10
			x /= 10
		}
		y = -y
	}

	if y > int(math.Exp2(float64(31))-1) || y < -int(math.Exp2(float64(31))) {
		return 0
	}
	return y
}

// 双百答案
func reverse(x int) int {
	ret := 0
	for ; x != 0; x /= 10 {
		ret = ret*10 + x%10
		if ret < math.MinInt32 || ret > math.MaxInt32 {
			return 0
		}
	}
	return ret
}

// 三数之和
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]]
// 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0
// 请你返回所有和为 0 且不重复的三元组
func threeSum_zhzw(nums []int) [][]int {
	N := len(nums)
	if N < 3 {
		return [][]int{}
	}

	var tmpResult [][]int
	var res []int
	for i := 0; i <= N-3; i++ {
		for j := i + 1; j <= N-3; j++ {
			for k := 0; k <= N-3; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]}...)
					tmpResult = append(tmpResult, res)
					res = []int{}
				}
			}
		}
	}
	// 二维切片去重，繁琐，此路不通
	var result [][]int
	return result
}

// 双指针方法
func threeSum(nums []int) [][]int {
	ans := [][]int{}
	N := len(nums)
	if nums == nil || N < 3 {
		return ans
	}
	// 先排序，保证唯一性，去重的前提
	sort.Ints(nums)
	// 此层for循环控制nums[i]的定位一直循环到nums[i] == 0
	for index, value := range nums {
		if value > 0 {
			break
		}
		// 排除值重复的固定位，去重的关键
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		// 指针初始位置，固定位右边第一个和数组最后一个
		left := index + 1
		right := N - 1
		for left < right {
			sum := value + nums[left] + nums[right]
			switch {
			case sum == 0:
				res := []int{value, nums[left], nums[right]}
				ans = append(ans, res)
				// 需要先行判断 left < right
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				// 准备下一次sum求值判断
				left++
				right--
			case sum > 0:
				// 和大于 0，说明 right 值大，左移
				right--
			case sum < 0:
				// 和小于 0，说明 left 值小，右移
				left++
			}
		}
	}
	return ans
}

// 最接近的三数之和
// 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
// 返回这三个数的和。
// 假定每组输入只存在恰好一个解。
// 暴力解法，超时，需用其他方法
func threeSumClosestZhzw(nums []int, target int) int {
	N := len(nums)
	var sumOfThree, sumToTarget, tmpSumToTarget int

	var res int
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := N - 1; k > j; k-- {
				sumOfThree = nums[i] + nums[j] + nums[k]
				tmpSumToTarget = int(math.Abs(float64(sumOfThree) - float64(target)))
				// 只初始化一次，后续满足条件才更新
				if i == 0 && j == 1 && k == N-1 || tmpSumToTarget < sumToTarget {
					sumToTarget = tmpSumToTarget
					res = sumOfThree
				}
			}
		}
	}
	return res
}

// 双指针法
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n    = len(nums)
		best = math.MaxInt32
	)

	// 枚举 a
	for i := 0; i < n; i++ {
		// 保证和上一次枚举的元素不相等
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 使用双指针枚举 b(j) 和 c(k)
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			// 如果和为 target 直接返回答案
			if sum == target {
				return target
			}
			// 根据差值的绝对值来更新答案
			if math.Abs(float64(sum-target)) < math.Abs(float64(best-target)) {
				best = sum
			}
			// 移动指针的优化
			if sum > target {
				// 如果和大于 target，移动 c 对应的指针
				nextK := k - 1
				// 跳过相等的元素
				for j < nextK && nums[nextK] == nums[k] {
					nextK--
				}
				k = nextK
			} else {
				// 如果和小于 target，移动 b 对应的指针
				nextJ := j + 1
				for nextJ < k && nums[nextJ] == nums[j] {
					nextJ++
				}
				j = nextJ
			}
		}
	}
	return best
}

// 四数之和
// 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。
// 请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]]
// （若两个四元组元素一一对应，则认为两个四元组重复）
func fourSum_zhzw(nums []int, target int) [][]int {
	N := len(nums)
	if N < 4 {
		return [][]int{}
	}
	// 先排序，处理不重复问题
	sort.Ints(nums)

	result := [][]int{}
	left, right := 0, N-1
	forsum := math.MaxInt64
	for i := 0; i < N; i++ {
		for left+1 < right-1 {
			midleft, midright := left+1, right-1
			forsum = nums[left] + nums[midleft] + nums[midright] + nums[right]
			if forsum == target {
				result = append(result, []int{nums[left], nums[midleft], nums[midright], nums[right]})
			} else if forsum > target {
				midright--
			} else {
				midleft++
			}

			// 内层遍历完
			if midleft >= midright {
				break
			}
		}
	}
	return result
}

// 四数之和，官方
func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	n := len(nums)
	// nums[i]+nums[i+1]+nums[i+2]+nums[i+3]>target即可退出循环
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		// nums[i] == nums[i-1]时会造成重复，找下一个i，枚举nums[i+1]
		// nums[i]+nums[n−3]+nums[n−2]+nums[n−1]<target时所有后续都不满足，找下一个i，枚举nums[i+1]
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		// 确定前两个数之后，如果nums[i]+nums[j]+nums[j+1]+nums[j+2]>target，四数之和一定大于target，退出循环
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			// nums[j] == nums[j-1]时会造成重复，找下一个j，枚举nums[j+1]
			// 确定前两个数之后，如果nums[i]+nums[j]+nums[n−2]+nums[n−1]<target，四数之和一定小于target，找下一个i，枚举nums[i+1]
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			// 剩下两个数用双指针
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					// 去除left重复项
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					// 去除right重复项
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return result
}
