package main

import "fmt"

func main() {
	// 1->2->3->3->4->4->5
	n5 := &ListNode{val: 5, next: nil}
	n42 := &ListNode{val: 4, next: n5}
	n41 := &ListNode{val: 4, next: n42}
	n33 := &ListNode{val: 3, next: n41}
	n32 := &ListNode{val: 3, next: n33}
	n31 := &ListNode{val: 3, next: n32}
	// n2 := &ListNode{val: 2, next: n31}
	// n1 := &ListNode{val: 1, next: n2}

	// for h := n1; h != nil; {
	// 	fmt.Printf("origin: %v \n", h.val)
	// 	h = h.next
	// }

	// deleteDuplicates(n1)
	// for h := n1; h != nil; {
	// 	fmt.Printf("%v ", h.val)
	// 	h = h.next
	// }

	for h := deleteDuplicates(n31); h != nil; {
		fmt.Printf("%v ", h.val)
		h = h.next
	}
}

// definition for singly-linked list.
type ListNode struct {
	val  int
	next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}

	h := &ListNode{next: head}
	f := h

	var val int
	if head.val == head.next.val {
		val = head.val
	} else {
		val = head.val - 1
	}

	for cur := head; ; {
		if cur.val == val {
			cur = cur.next
			continue
		}

		if n := cur.next; n != nil && cur.val == n.val {
			val = cur.val
			continue
		}

		val = cur.val
		f.next = cur
		f = cur
		cur = cur.next
		if cur == nil {
			f.next = nil
			break
		}
	}
	return h.next
}

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	res := make([][]int, 0, 10) // not correct
	for _, s := range subsets(nums[1:]) {
		res = append(res, s)
		res = append(res, append([][]int{{nums[0]}}, s)...)
	}
	return res
}

func reverse(x int) int {
	b := []byte(x)
}
