package utils

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkList(slice []int, length int) *ListNode {
	if length == 0 {
		return nil
	}
	head := &ListNode{Val: slice[0]}
	cur := head

	for i := 0; i < length; i++ {
		cur.Next = &ListNode{Val: slice[i]}
		cur = cur.Next
	}
	return head
}

func PrintLinklist(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("null\n")
}
