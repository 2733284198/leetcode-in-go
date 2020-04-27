package problem0096

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para int
	ans  int
}{
	{
		3,
		5,
	},
	{
		10,
		16796,
	},
}

func Test_NumTrees(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, numTrees(q.para), "输入：%v\n", q)
	}
}

func Test_NumTreesByDp(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, numTreesByDp(q.para), "输入：%v\n", q)
	}
}
