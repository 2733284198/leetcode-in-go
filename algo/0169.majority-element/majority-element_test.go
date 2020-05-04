package problem0169

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
		[]int{3, 2, 3},
		3,
	},
	{
		[]int{2, 2, 1, 1, 1, 2, 2},
		2,
	},
	{
		[]int{2, 2, 1, 1, 1},
		1,
	},
}

func Test_MarjorityElement(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, majorityElement(q.para), "输入：%v\n", q)
	}
}

func Test_BoyerMooreVote(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, boyerMooreVote(q.para), "输入：%v\n", q)
	}
}
