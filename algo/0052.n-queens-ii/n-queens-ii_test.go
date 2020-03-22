package problem0052

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
		4,
		2,
	},
	{
		5,
		10,
	},
}

func Test_totalNQueens(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, totalNQueens(q.para), "输入：%v\n", q)
	}
}
