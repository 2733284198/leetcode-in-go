package problem0230

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

	ans int
}{
	{
		[]int{2, 1, 3},
		[]int{1, 2, 3},
		1,
		1,
	},
	{
		[]int{3, 2, 1, 4, 5},
		[]int{1, 2, 3, 4, 5},
		3,
		3,
	},
}

func Test_KthSmallest(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, kthSmallest(kit.PreIn2Tree(q.para1, q.para2), q.para3), "输入：%v\n", q)
	}
}
