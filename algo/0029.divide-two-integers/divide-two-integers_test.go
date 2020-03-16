package problem0029

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 int
	para2 int
	ans   int
}{
	{
		10,
		3,
		3,
	},
	{
		7,
		-3,
		-2,
	},
}

func Test_Divide(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, divide(q.para1, q.para2), "输入：%v\n", q)
	}

}

func Benchmark_Divide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			divide(q.para1, q.para2)
		}
	}
}
