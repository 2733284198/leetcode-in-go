package problem0239

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 int
	ans   []int
}{
	{
		[]int{1, 3, -1, -3, 5, 3, 6, 7},
		3,
		[]int{3, 3, 5, 5, 6, 7},
	},
}

func Test_MaxSlidingWindow(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, maxSlidingWindow(q.para1, q.para2), "输入：%v\n", q)
	}
}
