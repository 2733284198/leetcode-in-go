package problem0062

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1, para2 int
	ans          int
}{
	{
		3, 2,
		3,
	},
	{
		7, 3,
		28,
	},
}

func Test_UniquePaths(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, uniquePaths(q.para1, q.para2), "输入：%v", q)
	}
}
