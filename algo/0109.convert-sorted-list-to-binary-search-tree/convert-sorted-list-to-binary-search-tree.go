package problem0109

import "github.com/aierui/leetcode-in-go/kit"

type TreeNode = kit.TreeNode
type ListNode = kit.ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	mid := findMidNode(head)

	node := &TreeNode{
		Val: mid.Val,
	}

	if head == mid {
		return node
	}

	node.Left = sortedListToBST(head)
	node.Right = sortedListToBST(mid.Next)

	return node
}

func findMidNode(head *ListNode) *ListNode {

	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		prev = slow
	}

	if prev != nil {
		prev.Next = nil
	}

	return slow
}
