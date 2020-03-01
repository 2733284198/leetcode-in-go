package problem0104

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

func Test_MaxDepth(t *testing.T) {
	ast := assert.New(t)
	qs := []struct {
		pre, in []int
		ans     int
	}{
		{
			[]int{},
			[]int{},
			0,
		},
		{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			3,
		},
		{
			[]int{3, 9, 20, 15, 10, 7},
			[]int{9, 3, 10, 15, 20, 7},
			4,
		},
	}

	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, maxDepth(kit.PreIn2Tree(q.pre, q.in)), "输入:%v", q)
	}
}
