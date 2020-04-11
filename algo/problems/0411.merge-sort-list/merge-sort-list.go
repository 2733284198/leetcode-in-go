package problems0411

// ListNode singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeSortLinkedList(list *ListNode) *ListNode {
	if list == nil {
		return nil
	}

	if list.Next == nil {
		return list
	}

	midList := findMidList(list)

	return merge(mergeSortLinkedList(list), mergeSortLinkedList(midList))
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	prev := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			head.Next = l2
			l2 = l2.Next
		} else {
			head.Next = l1
			l1 = l1.Next
		}
		head = head.Next
	}

	if l1 == nil {
		head.Next = l2
	} else {
		head.Next = l1
	}

	return prev.Next
}

func findMidList(list *ListNode) *ListNode {
	fast, slow := list, list

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	midList := slow.Next
	slow.Next = nil

	return midList
}
