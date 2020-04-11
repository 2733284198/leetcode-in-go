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

var qs = []struct {
	para1 []int
	para2 int
	para3 int
	ans   []int
}{

	//{
	//	[]int{1, 2, 3},
	//	1,
	//	2,
	//	[]int{2, 1, 3},
	//},
	{
		[]int{1, 2, 3, 4, 5},
		2,
		4,
		[]int{1, 4, 3, 2, 5},
	},
}

func Test_ReverseBetween(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, l2s(reverseBetween(s2l(q.para1), q.para2, q.para3)), "输入:%v", q)
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
