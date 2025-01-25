package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	head := &ListNode{}
	curr := head

	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next

		// carry if necessary, reset sum otherwise
		if sum >= 10 {
			sum = 1
		} else {
			sum = 0
		}
	}

	// if we get to the end and still have carry, append to the end
	if sum > 0 {
		curr.Next = &ListNode{Val: sum}
	}

	return head.Next
}
