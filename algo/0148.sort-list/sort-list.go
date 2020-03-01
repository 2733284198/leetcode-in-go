package problem0148

import (
	"github.com/aierui/leetcode-in-go/kit"
)

// ListNode export singly-linked list
type ListNode = kit.ListNode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	n := slow.Next
	slow.Next = nil

	return merge(sortList(head), sortList(n))
}

func merge(node1, node2 *ListNode) *ListNode {
	node := &ListNode{Val: 0}
	cur := node

	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			cur.Next, node1 = node1, node1.Next
		} else {
			cur.Next, node2 = node2, node2.Next
		}
		cur = cur.Next
	}
	if node1 != nil {
		cur.Next, node1 = node1, node1.Next
	}
	if node2 != nil {
		cur.Next, node2 = node2, node2.Next
	}

	return node.Next
}
