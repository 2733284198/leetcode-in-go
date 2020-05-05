package problem0019

import "github.com/aierui/leetcode-in-go/kit"

type ListNode = kit.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{Val: 0, Next: head} // 哨兵节点
	var prev *ListNode
	cur := result
	i := 1

	for head != nil {
		if i >= n {
			prev = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}

	prev.Next = prev.Next.Next

	return result.Next
}

func removeNthFronEndV2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	fast, slow := dummy, dummy

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
