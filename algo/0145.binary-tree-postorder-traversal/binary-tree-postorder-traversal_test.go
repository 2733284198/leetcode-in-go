package problem0145

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 []int
	ans   []int
}{
	{
		[]int{3, 9, 20, 15, 7},
		[]int{9, 3, 15, 20, 7},
		[]int{9, 15, 7, 20, 3},
	},
}

func Test_BinaryTreePostOrderTraversal(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, postorderTraversal(kit.PreIn2Tree(q.para1, q.para2)), "输入：%v\n", q)
	}
}
