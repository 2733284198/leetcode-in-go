package problem0065

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para string
	ans  bool
}{
	{
		"0",
		true,
	},
	{
		" 0.1",
		true,
	},
	{
		"abc",
		false,
	},
	{
		"1 a",
		false,
	},
	{
		"2e10",
		true,
	},
	{
		"95a54e53",
		false,
	},
	{
		" 1e",
		false,
	},
	{
		"-+3",
		false,
	},
	{
		" --6 ",
		false,
	},
	{
		"53.5e93",
		true,
	},
}

func Test_IsNumber(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, isNumber(q.para), "输入：%v\n", q)
	}
}
