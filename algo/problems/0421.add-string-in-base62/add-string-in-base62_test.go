package problems0421

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 string
	para2 string
	ans   string
}{
	{
		"a0",
		"13",
		"b3",
	},
}

func Test_AddStringInBase62(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, addStringInBase62(q.para1, q.para2), "输入：%v\n", q)
	}
}
