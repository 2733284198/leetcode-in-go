package problem0118

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para int
	ans  [][]int
}{
	{
		5,
		[][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		},
	},
}

func Test_Generate(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		fmt.Printf("~~%v~~", q)
		ast.Equal(q.ans, generate(q.para), "输入：%v\n", q)
	}

}

func Test_GenerateByDp(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		fmt.Printf("~~%v~~", q)
		ast.Equal(q.ans, generateByDp(q.para), "输入：%v\n", q)
	}

}
