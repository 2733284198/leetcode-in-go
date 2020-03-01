package problem0237

import "github.com/aierui/leetcode-in-go/kit"

// ListNode export single-linked list
type ListNode = kit.ListNode

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
