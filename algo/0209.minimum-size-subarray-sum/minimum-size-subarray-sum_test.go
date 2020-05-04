package problem0209

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)


var qs =[]struct {
	para1 int
	para2 []int
	ans int
} {
	{
		7,
		[]int{2,3,1,2,4,3},
		2,
	},
}

func Test_MinSubArrayLen(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, minSubArrayLen(q.para1, q.para2), "输入：%v\n", q)
	}

}