package main

import (
	"link_list/utils"
)

// 翻转列表
func reverseList(head *utils.ListNode) *utils.ListNode {
	var pre *utils.ListNode
	var next *utils.ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {
	var a []int
	b := 25
	for i := 0; i < b; i++ {
		a = append(a, i)
	}

	link_list := utils.CreateLinkList(a, b)
	utils.PrintLinklist(link_list)

	reversed_link_list := reverseList(link_list)
	utils.PrintLinklist(reversed_link_list)
}
