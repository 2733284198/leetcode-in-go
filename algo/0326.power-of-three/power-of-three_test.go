package problem0326

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para int
	ans  bool
}{
	{
		27,
		true,
	},
	{
		0,
		false,
	},
	{
		9,
		true,
	},
	{
		45,
		false,
	},
}

func Test_IsPowerOfThree(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, isPowerOfThree(q.para), "输入：%v\n", q)
	}
}

func Test_IsPowerOfThreeByMaxInt(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, isPowerOfThreeByMaxInt(q.para), "输入：%v\n", q)
	}
}
