package problem0557

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para string
	ans  string
}{

	{
		"Let's take LeetCode contest",
		"s'teL ekat edoCteeL tsetnoc",
	},
}

func Test_ReverseWords(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, reverseWords(q.para), "输入：%v\n", q)
	}
}
