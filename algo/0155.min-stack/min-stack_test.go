package problem0155

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	ast := assert.New(t)

	s := Constructor()

	s.Push(-2)
	// [-2]
	s.Push(0)
	// [-2, 0]
	s.Push(-3)
	// [-2, 0, -3]
	ast.Equal(-3, s.GetMin(), "get minStack from [-2, 0, -3]")
	// [-2, 0, -3]
	s.Pop()
	// [-2, 0]
	ast.Equal(0, s.Top(), "get top from [-2, 0]")
	// [-2, 0]
	ast.Equal(-2, s.GetMin(), "get minStack from [-2, 0]")
	// [-2, 0]
}
