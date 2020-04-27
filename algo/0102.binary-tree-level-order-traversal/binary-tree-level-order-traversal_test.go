package problem0102

import (
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)


var qs =[]struct {
	para1 []int
	para2 []int
	ans [][]int
} {
	{
		[]int{}
	}
}

func Test_LevelOrder(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		root := kit.PreIn2Tree(q.para1, q.para2)
		levelOrder(root)
		ast.Equal(q.ans, root, "输入：%v\n", q)
	}
}
