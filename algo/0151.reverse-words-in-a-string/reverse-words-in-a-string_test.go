package problem0151

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
		"hello world",
		"world hello",
	},
	{
		"the sky is blue",
		"blue is sky the",
	},
}

func Test_ReverseWords(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, reverseWords(q.para), "输入：%v\n", q)
	}
}

func Test_ReverseWordsV2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, reverseWordsV2(q.para), "输入：%v\n", q)
	}
}
