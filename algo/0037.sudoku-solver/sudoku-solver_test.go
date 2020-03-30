package problem0037

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para [][]byte
	ans  [][]byte
}{
	{
		[][]byte{
			[]byte(".87654321"),
			[]byte("2........"),
			[]byte("3........"),
			[]byte("4........"),
			[]byte("5........"),
			[]byte("6........"),
			[]byte("7........"),
			[]byte("8........"),
			[]byte("9........"),
		},
		[][]byte{
			[]byte("519748632"),
			[]byte("783652419"),
			[]byte("426139875"),
			[]byte("357986241"),
			[]byte("264317598"),
			[]byte("198524367"),
			[]byte("975863124"),
			[]byte("832491756"),
			[]byte("641275983"),
		},
	},
}

func Test_IsValidSudoku(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		solveSudoku(q.para)
		ast.Equal(q.para, q.para, "输入：%v\n", q)
	}
}
