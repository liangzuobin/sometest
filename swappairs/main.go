package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var former *ListNode
	for cur := head; cur != nil && cur.Next != nil; {
		next := cur.Next
		cur.Next, next.Next = next.Next, cur
		if former != nil {
			former.Next = next
		} else {
			head = next
		}
		former, cur = cur, cur.Next
	}
	return head
}

func main() {
	a := &ListNode{Val: 1, Next: nil}
	b := &ListNode{Val: 2, Next: a}
	c := &ListNode{Val: 3, Next: b}
	for n := c; n != nil; n = n.Next {
		fmt.Printf("original: (%v, %+v) \n", n.Val, n.Next)
	}
	for n := swapPairs(c); n != nil; n = n.Next {
		fmt.Printf("swaped: (%v, %+v) \n", n.Val, n.Next)
	}
}
