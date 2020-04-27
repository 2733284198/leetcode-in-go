package problem0107

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 []int
	ans   [][]int
}{
	{
		[]int{3, 9, 20, 15, 7},
		[]int{9, 3, 15, 20, 7},
		[][]int{
			{15, 7},
			{9, 20},
			{3},
		},
	},
}

func Test_LevelOrderBOttom(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		root := kit.PreIn2Tree(q.para1, q.para2)
		ast.Equal(q.ans, levelOrderBottom(root), "输入：%v\n", q)
	}
}
