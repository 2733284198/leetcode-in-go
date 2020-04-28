package problem0110

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 []int
	ans   bool
}{
	{
		[]int{3, 9, 20, 15, 7},
		[]int{9, 3, 15, 20, 7},
		true,
	},
	{
		[]int{1, 2, 3, 4, 4, 3, 2},
		[]int{4, 3, 4, 2, 3, 1, 2},
		false,
	},
}

func Test_IsBalanced(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, isBalanced(kit.PreIn2Tree(q.para1, q.para2)), "输入：%v\n", q)
	}
}
