package problem0226

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 []int
	ans1  []int
	ans2  []int
}{
	{
		[]int{4, 2, 1, 3, 7, 6, 9},
		[]int{1, 2, 3, 4, 5, 7, 9},
		[]int{4, 7, 9, 6, 2, 3, 1},
		[]int{9, 7, 6, 4, 3, 2, 1},
	},
}

func Test_InvertTree(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(kit.PreIn2Tree(q.ans1, q.ans2), invertTree(kit.PreIn2Tree(q.para1, q.para2)))
	}
}
