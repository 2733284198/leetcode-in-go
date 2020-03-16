package problem0092

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
	one   []int
	two   int
	three int
}

type ans struct {
	one []int
}

func Test_Problem0092(t *testing.T) {

	ast := assert.New(t)

	qs := []question{
		question{
			para{
				[]int{1, 2, 3},
				1,
				2,
			},
			ans{
				[]int{2, 1, 3},
			},
		},
		question{
			para{
				[]int{1, 2, 3, 4, 5},
				2,
				4,
			},
			ans{
				[]int{1, 4, 3, 2, 5},
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		ast.Equal(a.one, l2s(reverseBetween(s2l(p.one), p.two, p.three)), "输入:%v", p)
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
