package problem0025

import "github.com/aierui/leetcode-in-go/kit"

// ListNode Definition for singly-linked list.
type ListNode = kit.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	tail := head
	for i := 0; i < k-1; i++ {
		if tail == nil {
			return head
		}
		tail = tail.Next
	}

	if tail == nil {
		return head
	}

	var pre, cur, next *ListNode

	pre = reverseKGroup(tail.Next, k)

	tail.Next = nil
	cur = head
	next = cur.Next

	for {
		cur.Next = pre
		pre = cur

		if next == nil {
			break
		}

		cur = next
		next = cur.Next
	}

	return cur
}
