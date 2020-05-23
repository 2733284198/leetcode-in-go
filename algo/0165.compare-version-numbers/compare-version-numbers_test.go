package problem0165

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var qs = []struct {
	para1, para2 string
	ans          int
}{
	{
		"0.1",
		"1.1",
		-1,
	},
	{
		"1.0.1",
		"1",
		1,
	},
	{
		"7.5.2.4",
		"7.5.3",
		-1,
	},
	{
		"1.01",
		"1.001",
		0,
	},
	{
		"1.0",
		"1.0.0",
		0,
	},
}

func Test_CompareVersion(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, compareVersion(q.para1, q.para2), "输入：%v\n", q)
	}
}
