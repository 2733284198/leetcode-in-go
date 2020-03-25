package problem0069

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para int
	ans  int
}{
	{
		4,
		2,
	},
	{
		8,
		2,
	},
}

func Test_MySqrt(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, mySqrt(q.para), "输入：%v\n", q)
	}
}
