package problem0043

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 string
	para2 string

	ans string
}{
	{
		"3",
		"2",
		"6",
	},
	{
		"123",
		"456",
		"56088",
	},
	{
		"1234567890",
		"4561234567234",
		"5631153735465142516260",
	},
	{
		"12345678905632341153343273546514243516260",
		"45612345672345631153234735465142516260",
		"563115373803488063666544004426994190080372273560949507970698672183566624387600",
	},
}

func Test_Multiply(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, multiply(q.para1, q.para2), "输入：%v\n", q)
	}
}
