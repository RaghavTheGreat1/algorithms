package linked_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	header := &ListNode{Next: head}
	temp := head
	prev := header

	count := 0

	// Count total number of nodes
	for temp != nil {
		temp = temp.Next
		count++
	}

	// Reset temp for re-use
	temp = head
	count = count - n

	// Iterate through count - n from front
	for temp != nil && count != 0 {
		prev = temp
		temp = temp.Next
		count--
	}

	// Remove element
	prev.Next = temp.Next

	return header.Next
}
