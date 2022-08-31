package main

import (
	"fmt"
	"link_list/utils"
)

// 删除指定值的节点
func removeElements() {
	fmt.Println()
}

// 翻转列表
func reverseList(head *utils.ListNode) *utils.ListNode {
	var pre, next *utils.ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur  // 循环结束时，pre为非nil的最后一个节点，即头节点
		cur = next // 循环结束时，cur为nil
	}
	return pre
}

func main() {
	var a []int
	b := 10
	for i := 0; i < b; i++ {
		a = append(a, i)
	}

	link_list := utils.CreateLinkList(a, b)
	utils.PrintLinklist(link_list)
	// reversed_link_list := reverseList(link_list)
	// utils.PrintLinklist(reversed_link_list)
}
