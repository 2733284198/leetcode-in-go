package problem0098

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
		[]int{2, 1, 3},
		[]int{1, 2, 3},
		true,
	},
	{
		[]int{5, 1, 4, 3, 6},
		[]int{1, 5, 3, 4, 6},
		false,
	},
}

func Test(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, isValidBST(kit.PreIn2Tree(q.para1, q.para2)), "输入：%v\n", q)
	}
}
