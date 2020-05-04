package problem0160

import "github.com/aierui/leetcode-in-go/kit"

type ListNode = kit.ListNode

func getInterscetionNode(headA, headB *ListNode) *ListNode {
	pA := headA
	pB := headB

	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA
}