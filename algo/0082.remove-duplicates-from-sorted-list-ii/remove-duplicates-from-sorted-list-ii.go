package problem0082

import "github.com/aierui/leetcode-in-go/kit"

// ListNode defined single linked list
type ListNode = kit.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	prev := &ListNode{0, head}
	cur := prev
	count := 0

	for head != nil && head.Next != nil {

		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
			count++
		}

		if count > 0 {
			cur.Next = head.Next
			head = head.Next
			count = 0
		} else {
			cur = head
			head = head.Next
		}
	}

	return prev.Next
}
