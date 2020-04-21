package problem0128

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para []int
	ans  int
}{
	{
		[]int{100, 4, 200, 1, 3, 2},
		4,
	},
	{
		[]int{100, 5, 200, 1, 3, 2},
		3,
	},
	{
		[]int{100, 5, 200, 1, 3, 2, 4},
		5,
	},
}

func Test_LongestConsecutive(t *testing.T) {

	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, longestConsecutive(q.para), "输入：%v\n", q)
	}
}
