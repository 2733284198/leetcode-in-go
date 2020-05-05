package problem0019

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 int
	ans   []int
}{
	{
		[]int{1, 2, 3, 4, 5},
		2,
		[]int{1, 2, 3, 5},
	},
	{
		[]int{3, 4, 1, 2, 3, 4, 5},
		5,
		[]int{3, 4, 2, 3, 4, 5},
	},
}

func Test_RemoveNthNodeFromEndOfList(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, kit.List2Ints(removeNthFromEnd(kit.Ints2List(q.para1), q.para2)), "输入：%v\n", q)
	}
}

func Test_RemoveNthNodeFromEndV2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, kit.List2Ints(removeNthFronEndV2(kit.Ints2List(q.para1), q.para2)), "输入：%v\n", q)
	}
}
