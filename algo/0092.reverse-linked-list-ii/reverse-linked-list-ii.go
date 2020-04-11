package problem0092

// ListNode export single-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m, n int) *ListNode {
	if head == nil || m == n {
		return head
	}

	prev, i, j := &ListNode{Next: head}, m, n-m
	head = prev

	for i > 1 {
		prev = prev.Next
		i--
	}

	cur := prev.Next // m 个节点

	for j > 0 {

		temp := cur.Next      // 保存要反转节点的下一个节点
		cur.Next = temp.Next  // 当前节点指向 要放转节点的 next 节点 最终指向第 n + 1 个节点
		temp.Next = prev.Next // 反转
		prev.Next = temp      // 第 m 个节点指向后面一个节点

		j--
	}

	return head.Next
}
