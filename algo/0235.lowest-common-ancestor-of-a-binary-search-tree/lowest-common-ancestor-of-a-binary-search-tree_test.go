package problem0235

import (
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1, para2 []int
	p, q, ans    int
}{
	{
		[]int{6, 2, 0, 4, 3, 5, 8, 7, 9},
		[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
		2,
		8,
		6,
	},
}

func Test_LowestCommonAncestor(t *testing.T) {
	ast := assert.New(t)
	for _, item := range qs {
		root := kit.PreIn2Tree(item.para1, item.para2)

		p := &TreeNode{
			Val: item.p,
		}

		q := &TreeNode{
			Val: item.q,
		}

		node := lowestCommonAncestor(root, p, q)

		ast.Equal(item.ans, node.Val)
	}
}
