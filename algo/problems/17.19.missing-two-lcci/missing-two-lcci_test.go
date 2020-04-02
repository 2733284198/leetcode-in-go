package problems1719

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
		[]int{1},
		[]int{2, 3},
	},
	{
		[]int{2, 3},
		[]int{1, 4},
	},
}

func Test_MissingTwo(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, missingTwo(q.para), "输入：%v", q)
	}
}
