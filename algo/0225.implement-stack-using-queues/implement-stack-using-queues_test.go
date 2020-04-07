package problem0225

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	ast := assert.New(t)

	s := Constructor()
	ast.True(s.Empty(), "检查新建的 s 是否为空")

	start, end := 0, 10

	for i := start; i < end; i++ {
		s.Push(i)
		ast.Equal(i, s.Top(), "查看 s.Top()")
	}

	for i := end - 1; i >= start; i-- {
		ast.Equal(i, s.Pop(), "从 s 中 pop 出数来。")
	}

	ast.True(s.Empty(), "检查 Pop 完毕后的 s 是否为空")
}
