package problem0083

import "github.com/aierui/leetcode-in-go/kit"

// ListNode
type ListNode = kit.ListNode

// func deleteDuplicates(head *ListNode) *ListNode  {
// 	cur := head
// 	for cur != nil {
// 		for cur.Next != nil && cur.Next.Val == cur.Val {
// 			cur.Next = cur.Next.Next
// 		}
// 		cur = cur.Next
// 	}

// 	return head
// }

func deleteDuplicates(head *ListNode) *ListNode  {
	if head == nil {
		return head
	}

	// 定义两个指针
	prev, cur := head, head.Next 

	for cur != nil {
		if cur.Val == prev.Val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}

	return head
}