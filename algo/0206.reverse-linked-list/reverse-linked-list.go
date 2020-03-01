package problem0206

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList iterative func
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // 二者的顺序不可调换
		return head
	}
	var prev *ListNode

	// head 是下一个被逆转的节点
	for head != nil {
		// 让temp指向head.Next, 临时保存起来后面用
		temp := head.Next
		// head称为已经逆转的节点的新head
		head.Next = prev
		prev = head
		// 让head指向下一个被逆转的节点
		head = temp
	}

	return prev
}

// reverseListByRecursion recursion func
func reverseListByRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // 二者的顺序不可调换
		return head
	}
	last := reverseListByRecursion(head.Next)

	head.Next.Next = head
	head.Next = nil

	return last
}
