package problem0028

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 string
	para2 string
	ans   int
}{

	{
		"hello",
		"ll",
		2,
	},
	{
		"aaaaa",
		"bba",
		-1,
	},
}

func Test_Strstr(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, strStr(q.para1, q.para2), "输入：%v\n", q)
	}
}

func Test_StrStrNotUseBuiltInFunc(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, strStrNotUseBuiltInFunc(q.para1, q.para2), "输入：%v\n", q)
	}
}
