package arrays

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	l3 := result

	for l1 != nil || l2 != nil {
		if l1 != nil {
			l3.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l3.Val += l2.Val
			l2 = l2.Next
		}
		if l3.Val > 9 {
			l3.Val -= 10
			l3.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			l3.Next = &ListNode{}
		}
		l3 = l3.Next
	}

	return result
}
