package problem0234

import (
	"github.com/aierui/leetcode-in-go/kit"
)

type ListNode = kit.ListNode

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 1、找链表的中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	var prev *ListNode

	// 2、反转链表的尾部
	for slow != nil {
		temp := slow.Next
		slow.Next = prev
		prev = slow
		slow = temp
	}

	// 3、遍历尾部和头部对比
	for prev != nil {
		if prev.Val != head.Val {
			return false
		}
		prev = prev.Next
		head = head.Next
	}

	return true
}
