package problem0051

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para int
	ans  [][]string
}{
	{
		4,
		[][]string{
			[]string{".Q..", "...Q", "Q...", "..Q."},
			[]string{"..Q.", "Q...", "...Q", ".Q.."},
		},
	},
	{
		5,
		[][]string{
			[]string{"Q....", "..Q..", "....Q", ".Q...", "...Q."},
			[]string{"Q....", "...Q.", ".Q...", "....Q", "..Q.."},
			[]string{".Q...", "...Q.", "Q....", "..Q..", "....Q"},
			[]string{".Q...", "....Q", "..Q..", "Q....", "...Q."},
			[]string{"..Q..", "Q....", "...Q.", ".Q...", "....Q"},
			[]string{"..Q..", "....Q", ".Q...", "...Q.", "Q...."},
			[]string{"...Q.", "Q....", "..Q..", "....Q", ".Q..."},
			[]string{"...Q.", ".Q...", "....Q", "..Q..", "Q...."},
			[]string{"....Q", ".Q...", "...Q.", "Q....", "..Q.."},
			[]string{"....Q", "..Q..", "Q....", "...Q.", ".Q..."},
		},
	},
}

func Test_solveNQueens(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)

		result := solveNQueens(q.para)

		ast.Equal(len(q.ans), len(result), "input:%v\n", q)

		for i, v := range result {
			ast.Equal(q.ans[i], v, "input:%v\n", q)
		}
	}

}
