package problem0023

import "container/heap"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// PQ lists
type PQ []*ListNode

func (p PQ) Len() int { return len(p) }

func (p PQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PQ) Less(i, j int) bool {
	return p[i].Val < p[j].Val
}

// Push PQ
func (p *PQ) Push(x interface{}) {
	node := x.(*ListNode)
	*p = append(*p, node)
}

// Pop PQ
func (p *PQ) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListNode{
		Val:  -1,
		Next: nil,
	}
	t := h
	if len(lists) == 0 {
		return h.Next
	}

	pq := make(PQ, 0)
	for i := range lists {
		if lists[i] != nil {
			pq = append(pq, lists[i])
		}
	}
	heap.Init(&pq)

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*ListNode)
		next := item.Next

		item.Next = t.Next
		t.Next = item
		t = item

		if next != nil {
			heap.Push(&pq, next)
		}
	}

	return h.Next
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
