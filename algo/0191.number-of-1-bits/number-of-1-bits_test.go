package problem0191

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para uint32
	ans  int
}{
	{

		1,
		1,
	},
	{

		11,
		3,
	},
	{

		1024,
		1,
	},
}

func Test_HammingWeight(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, hammingWeight(q.para), "输入：%v", q)
	}
}
