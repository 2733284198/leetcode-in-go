package problem0142

import "github.com/aierui/leetcode-in-go/kit"

type ListNode = kit.ListNode

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	//fast = head
	for fast != head {
		fast = fast.Next
		head = head.Next
	}
	return fast
}
