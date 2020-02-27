package problem0061

// ListNode export singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 0 {
		return head
	}

	tail := head
	m := 1

	for tail.Next != nil { // 遍历一次链表计算长度
		tail = tail.Next
		m++
	}

	if k%m == 0 { // 不移动
		return head
	}

	tail.Next = head               // 让 tail 的 Next 指向头节点，形成一个环形链表
	for i := 0; i < m-k%m-1; i++ { // 遍历指向新的结束节点 （k %m  当 k 大于m时候）
		head = head.Next
	}

	result := head.Next // head 下一个节点作为新的头节点
	head.Next = nil     // head

	return result
}
