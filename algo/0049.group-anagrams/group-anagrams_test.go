package problem0049

import (
	"fmt"
	"sort"
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
			[]string{"bat"},
			[]string{"nat", "tan"},
			[]string{"ate", "eat", "tea"}},
	},
}

func Test_GroupAnagrams(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		result := groupAnagrams(q.para)
		for _, v := range result {
			sort.Strings(v)
			ast.Equal(q.ans[len(v)-1], v, "输入：%v\n", q)
		}
	}
}
