package problem0137

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
		[]int{2, 2, 3, 2},
		3,
	},
	{
		[]int{0, 1, 0, 1, 1, 2, 0, 99, 2, 2},
		99,
	},
}

func Test_SingleNumberHashMap(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, singleNumberHashMap(q.para), "输入：%v\n", q)
	}
}

func Test_SingleNumber(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, singleNumber(q.para), "输入：%v\n", q)
	}
}
