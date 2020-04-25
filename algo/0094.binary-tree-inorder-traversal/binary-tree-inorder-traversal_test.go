package problem0094

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
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		[]int{1, 3, 2},
	},
}

func Test_InOrderTraversalByRecursion(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, inOrderTraversalByRecursion(kit.PreIn2Tree(q.para1, q.para2)), "输入：%v\n", q)
	}
}
