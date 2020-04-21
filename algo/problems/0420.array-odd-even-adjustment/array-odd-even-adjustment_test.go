package problems0420

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
		[]int{1, 2, 3, 4, 5, 6},
		[]int{1, 3, 5, 2, 4, 6},
	},
}

func Test_ArrayAdjust(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		arrayAdjust(q.para)
		ast.Equal(q.ans, q.para, "输入：%v\n", q)
	}
}
