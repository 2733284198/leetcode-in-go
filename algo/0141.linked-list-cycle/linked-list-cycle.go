package problem0141

import "github.com/aierui/leetcode-in-go/kit"

// ListNode
type ListNode = kit.ListNode

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head
	slow := head
	for fast != nil && slow != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
