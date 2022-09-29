package utils

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkList(slice []int) *ListNode {
	length := len(slice)
	if length == 0 {
		return nil
	}
	head := &ListNode{Val: slice[0]}
	cur := head

	for i := 1; i < length; i++ {
		cur.Next = &ListNode{Val: slice[i]}
		cur = cur.Next
	}
	return head
}

func PrintLinklist(head *ListNode) {
	cur := head
	fmt.Printf("Head node's Val: %d\n", cur.Val)
	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		if cur.Next == nil {
			fmt.Printf("NULL\n")
			fmt.Printf("Tail node's Val: %d\n", cur.Val)
		}
		cur = cur.Next
	}
	fmt.Println("=========================")
}
