package problem0023

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

}

func mergeKListsFail(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	m := lists[0]
	n := mergeKListsFail(lists[:1])

	k := &ListNode{Val: -1}
	p := k

	for m != nil && n != nil {
		if m.Val < n.Val {
			p.Next = m
			m = m.Next
		} else {
			p.Next = n
			n = n.Next
		}
		p = p.Next
	}

	if m != nil {
		p.Next = m
	}

	if n != nil {
		p.Next = n
	}

	return k.Next
}
