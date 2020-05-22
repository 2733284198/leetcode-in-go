package problem0119

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para int
	ans  []int
}{
	{
		3,
		[]int{1, 3, 3, 1},
	},
}

func Test_GetRow(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		fmt.Printf("~~%v~~", q)
		ast.Equal(q.ans, getRow(q.para), "输入：%v\n", q)
	}

}

func Test_GetRowByMap(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		fmt.Printf("~~%v~~", q)
		ast.Equal(q.ans, getRowByMap(q.para), "输入：%v\n", q)
	}

}
