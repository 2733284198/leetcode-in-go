package problem0108


import (
	"testing"
	"fmt"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para []int
	ans []int
} {
	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{4, 2, 1, 3, 6, 5, 7},
	},
	{
		[]int{},
		nil,
	},
}

func Test_SortedArrayToBST(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, kit.Tree2Preorder(sortedArrayToBST(q.para)), "输入：%v\n", q)
	}
}