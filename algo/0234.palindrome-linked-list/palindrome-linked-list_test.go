package problem0234

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para []int
	ans  bool
}{
	{
		[]int{1, 2},
		false,
	},
	{
		[]int{1, 2, 2, 1},
		true,
	},
}

func Test_IsPalindrome(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, isPalindrome(kit.Ints2List(q.para)), "输入：%v\n", q)
	}
}
