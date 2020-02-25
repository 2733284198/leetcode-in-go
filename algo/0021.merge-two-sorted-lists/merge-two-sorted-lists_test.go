package problem0021

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	one []int
	two []int
}

type ans struct {
	one []int
}

func Test_Problem0021(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{[]int{1, 2, 4},
				[]int{1, 3, 4},
			},
			ans{[]int{1, 1, 2, 3, 4, 4}},
		},
		question{
			para{[]int{1, 3, 5, 7, 9},
				[]int{2, 4, 6},
			},
			ans{[]int{1, 2, 3, 4, 5, 6, 7, 9}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		ast.Equal(a.one, l2s(mergeTwoLists(s2l(p.one), s2l(p.two))), "输入:%v", p)
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
