package problem0058

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para string
	ans  int
}{
	{
		"Hello World",
		5,
	},
	{
		"World Hello",
		5,
	},
}

func Test_LengthOfLastWord(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, lengthOfLastWord(q.para), "输入：%v\n", q)
	}
}

func Test_LengthOfLastWordUseLibV2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, lengthOfLastWordUseLibV2(q.para), "输入：%v\n", q)
	}
}

func Test_LengthOfLastWordNotUseLib(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, lengthOfLastWordNotUseLib(q.para), "输入：%v\n", q)
	}
}

func Benchmark_LengthOfLastWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			lengthOfLastWord(q.para)
		}
	}
}

func Benchmark_LengthOfLastWordUseLibV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			lengthOfLastWordUseLibV2(q.para)
		}
	}
}

func Benchmark_LengthOfLastWordNotUseLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			lengthOfLastWordNotUseLib(q.para) // 不实用类库函数效率高的多
		}
	}
}
