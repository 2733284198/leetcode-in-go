package problem0013

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 string
	para2 []string
	ans   []int
}{
	{
		"barfoothefoobarman",
		[]string{"foo", "bar"},
		[]int{0, 9},
	},
	{
		"wordgoodgoodgoodbestword",
		[]string{"word", "good", "best", "word"},
		[]int{},
	},
}

func Test_FindSubString(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, findSubstring(q.para1, q.para2), "输入：%v\n", q)
	}
}
