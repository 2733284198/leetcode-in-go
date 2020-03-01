package problem0344

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	s   string
	ans string
}{
	{"hello", "olleh"},
	{"world", "dlrow"},
}

func Test_ReverseString(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, reverseString(q.s), "输入：%v \n", q)
	}
}

func Benchmark_ReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			reverseString(q.s)
		}
	}
}
