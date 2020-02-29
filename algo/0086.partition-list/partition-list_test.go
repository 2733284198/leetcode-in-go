package problem0086

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one []int
}

type question struct {
	para
	ans
}

func Test_Problem0086(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		question{
			para{
				[]int{1, 4, 3, 2, 5, 2},
				3,
			},
			ans{
				[]int{1, 2, 2, 4, 3, 5},
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		ast.Equal(a.one, l2s(partition(s2l(p.one), p.two)), "输入:%v", p)
	}
}

// convert *ListNode to []int
func l2s(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// convert []int to *ListNode
func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}
