package problem0017

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para string
	ans  []string
}{
	{
		"23",
		[]string{"ad", "bd", "cd", "ae", "be", "ce", "af", "bf", "cf"},
	},
}

func Test(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, letterCombinations(q.para), "输入：%v", q)
	}
}
