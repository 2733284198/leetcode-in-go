package problem0160

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	a, b []int
	m, n int
}{
	{
		[]int{4, 1, 8, 4, 5},
		[]int{5, 0, 1, 8, 4, 5},
		2,
		3,
	},
	{
		[]int{0, 9, 1, 2, 4},
		[]int{3, 2, 4},
		3,
		1,
	},
}

func Test_GetIntersctionNode(t *testing.T) {

	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)

		tail := kit.Ints2List(q.a[q.m:])

		headA := kit.Ints2List(q.a[:q.m])
		tailA := tailOf(headA)
		tailA.Next = tail

		headB := kit.Ints2List(q.b[:q.n])
		tailB := tailOf(headB)
		tailB.Next = tail

		ast.Equal(tail, getInterscetionNode(headA, headB), "输入：%v\n", q)
	}

}

func tailOf(head *ListNode) *ListNode {
	for head.Next != nil {
		head = head.Next
	}

	return head
}
