package problem0080

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
		[]int{1, 2, 3, 3, 3, 4, 4, 4},
		6,
	},
	{
		[]int{1, 2, 3, 4},
		4,
	},
}

func Test_RemoveDuplicates(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, removeDuplicates(q.para), "输入：%v\n", q)
	}
}

func Test_RemoveDuplicatesV2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, removeDuplicatesV2(q.para), "输入：%v\n", q)
	}
}
