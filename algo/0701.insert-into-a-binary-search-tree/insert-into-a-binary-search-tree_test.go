package problem0701

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 []int
	para3 int
	ans1  []int
	ans2  []int
}{
	{
		[]int{2, 1, 3},
		[]int{1, 2, 3},
		4,
		[]int{2, 1, 3, 4},
		[]int{1, 2, 3, 4},
	},
	{
		[]int{4, 2, 1, 3, 7},
		[]int{1, 2, 3, 4, 7},
		5,
		[]int{4, 2, 1, 3, 7, 5},
		[]int{1, 2, 3, 4, 5, 7},
	},
}

func Test_InsertIntoBST(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		root := kit.PreIn2Tree(q.para1, q.para2)
		tree := kit.PreIn2Tree(q.ans1, q.ans2)
		ast.Equal(tree, insertIntoBST(root, q.para3), "输入：%v\n", q)
	}
}
