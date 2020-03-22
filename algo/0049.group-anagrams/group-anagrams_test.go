package problem0049

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para []string
	ans  [][]string
}{
	{
		[]string{
			"eat", "tea", "tan", "ate", "nat", "bat",
		},
		[][]string{
			[]string{"eat", "tea", "ate"},
			[]string{"tan", "nat"},
			[]string{"bat"},
		},
	},
}

func Test_GroupAnagrams(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, groupAnagrams(q.para), "输入：%v\n", q)
	}
}
