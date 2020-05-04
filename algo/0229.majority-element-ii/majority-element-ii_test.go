package problem0229

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para []int
	ans  []int
}{
	{
		[]int{3, 2, 3},
		[]int{3},
	},
	{
		[]int{1, 1, 1, 3, 3, 2, 2, 2},
		[]int{1, 2},
	},
}

func Test_MajorityElement(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, majorityElement(q.para), "输入：%v\n", q)
	}
}
