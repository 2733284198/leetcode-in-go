package problem0231

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	one int
}

type ans struct {
	one bool
}

var qs = []question{
	question{
		para{
			1,
		},
		ans{
			true,
		},
	},
	question{
		para{
			16,
		},
		ans{
			true,
		},
	},
	question{
		para{
			218,
		},
		ans{
			false,
		},
	},
	question{
		para{
			0,
		},
		ans{
			false,
		},
	},
}

func Test_Problem0231(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		ast.Equal(a.one, isPowerOfTwo(p.one), "输入:%v", p)
	}
}
