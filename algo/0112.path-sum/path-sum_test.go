package problem0112

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
	ans   bool
}{
	{
		[]int{5, 4, 11, 7, 2, 8, 13, 4, 1},
		[]int{7, 11, 2, 4, 5, 13, 8, 4, 1},
		22,
		true,
	},
}

func Test_HasPathSum(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, hasPathSum(kit.PreIn2Tree(q.para1, q.para2), q.para3), "输入：%v\n", q)
	}
}