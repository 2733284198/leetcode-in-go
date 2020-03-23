package problem0055

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para []int
	ans  bool
}{
	{
		[]int{2, 3, 1, 1, 4},
		true,
	},
	{

		[]int{3, 2, 1, 0, 4},
		false,
	},
}

func Test_CanJump(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, canJump(q.para), "输入：%v\n", q)
	}
}

func Benchmark_CanJump(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
