package problem0099

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
		[]int{3, 2, 1},
		[]int{3, 2, 1},
		[]int{1, 2, 3},
	},
}

func Test_RecoverTree(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		root := kit.PreIn2Tree(q.para1, q.para2)
		recoverTree(root)
		ast.Equal(q.ans, kit.Tree2Inorder(root), "输入：%v\n", q)
	}
}
