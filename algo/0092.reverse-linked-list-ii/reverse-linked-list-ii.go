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

	d := prev

	for i > 1 {
		d = d.Next
		i--
	}

	cur := d.Next.Next
	curPre := d.Next

	for j > 0 {
		curPre.Next = cur.Next
		cur.Next = d.Next
		d.Next = cur
		cur = curPre.Next
		j--
	}

	return prev.Next
}
