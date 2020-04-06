package prolbem0125

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para string
	ans  bool
}{
	{
		"A man, a plan, a canal: Panama",
		true,
	},
	{
		"race a car",
		false,
	},
}

func Test_IsPalindrome(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, isPalindrome(q.para), "输入：%v\n", q)
	}
}
