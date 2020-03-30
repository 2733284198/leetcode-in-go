package problem0036

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para [][]byte
	ans bool
} {
	{
		[][]byte {
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
		true,
	},
	{
		[][]byte {
				[]byte(".87654321"),
				[]byte("2.3......"),
				[]byte("3........"),
				[]byte("4........"),
				[]byte("5........"),
				[]byte("6........"),
				[]byte("7........"),
				[]byte("8........"),
				[]byte("9........"),
		},
		false,
	},
	{
		[][]byte {
				[]byte(".88654321"),
				[]byte("2........"),
				[]byte("3........"),
				[]byte("4........"),
				[]byte("5........"),
				[]byte("6........"),
				[]byte("7........"),
				[]byte("8........"),
				[]byte("9........"),
		},
		false,
	},
}


func Test_IsValidSudoku(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, isValidSudoku(q.para), "输入：%v\n",q)
	}
}