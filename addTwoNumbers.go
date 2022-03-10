package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{0, nil}
	p := result

	sum := 0
	for {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		curr := sum
		if sum > 9 {
			curr = sum - 10
			sum = 1
		} else {
			sum = 0
		}

		p.Val = curr
		if l1 == nil && l2 == nil && sum == 0 {
			break
		}
		p.Next = &ListNode{0, nil}
		p = p.Next
	}

	return result
}
