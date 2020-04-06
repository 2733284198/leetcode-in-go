package problem0072

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
		"horse",
		"ros",
		3,
	},
	{
		"intention",
		"execution",
		5,
	},
}

func Test_MinDistance(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, minDistance(q.para1, q.para2), "输入：%v\n", q)
	}
}
