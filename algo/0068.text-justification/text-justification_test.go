package problem0068

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []string
	para2 int
	ans   []string
}{
	{
		[]string{"This", "is", "an", "example", "of", "text", "justification."},
		16,
		[]string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		},
	},
	{
		[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
		16,
		[]string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		},
	},
	{
		[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
		20,
		[]string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  ",
		},
	},
}

func Test_FullJustify(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, fullJustify(q.para1, q.para2), "输入：%v\n", q)
	}
}
