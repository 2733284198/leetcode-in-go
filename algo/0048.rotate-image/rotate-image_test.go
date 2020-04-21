package problem0048

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para [][]int
	ans  [][]int
}{
	{
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		},
		[][]int{
			[]int{7, 4, 1},
			[]int{8, 5, 2},
			[]int{9, 6, 3},
		},
	},
}

func Test_Rotate(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		rotate(q.para)
		ast.Equal(q.para, q.para, "输入：%v\n", q)
	}
}

func Test_RotateV2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		rotateV2(q.para)
		ast.Equal(q.para, q.para, "输入：%v\n", q)
	}
}
