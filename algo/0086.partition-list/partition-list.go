package problem0086

// ListNode export singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var prev, tail *ListNode

	prevNode, tailNode := prev, tail

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

	return prevNode.Next
}
