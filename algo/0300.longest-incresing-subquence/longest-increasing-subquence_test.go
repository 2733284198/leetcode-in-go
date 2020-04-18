package problem0300

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para []int
	ans  int
}{
	{
		[]int{10, 9, 2, 5, 3, 7, 101, 18},
		4,
	},
}

func Test_LengthOfLIS(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, lengthOfLIS(q.para), "输入：%v\n", q)
	}
}
