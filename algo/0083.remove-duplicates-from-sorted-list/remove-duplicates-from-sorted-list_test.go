package problem0083

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para []int
	ans  []int
}{
	{
		[]int{1, 1, 2, 3, 3},
		[]int{1, 2, 3},
	},
}

func Test_DeleteDuplicates(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, kit.List2Ints(deleteDuplicates(kit.Ints2List(q.para))), "输入：%v\n", q)
	}
}
