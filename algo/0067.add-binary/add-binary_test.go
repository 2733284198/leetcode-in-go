package problem0067

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1, para2 string
	ans          string
}{
	{
		"11",
		"1",
		"100",
	},
	{
		"1010",
		"1011",
		"10101",
	},
}

func Test_AddBinary(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, addBinary(q.para1, q.para2), "输入：%v\n", q)
	}
}
