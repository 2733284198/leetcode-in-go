package problem0086

// ListNode export singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	prev := &ListNode{
		Val:  -1,
		Next: nil,
	}

	tail := &ListNode{
		Val:  -1,
		Next: nil,
	}

	prevNode := prev
	tailNode := tail

	for head != nil {
		if head.Val < x {
			prevNode.Next = head
			prevNode = prevNode.Next
		} else {
			tailNode.Next = head
			tailNode = tailNode.Next
		}
		head = head.Next
	}

	tailNode.Next = nil
	prevNode.Next = tail.Next

	return prev.Next
}
