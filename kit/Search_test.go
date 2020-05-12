package kit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var searchQs = []struct {
	para1 []int
	para2 int
	ans   int
}{
	{
		[]int{2, 3, 32, 6, 4, 1, 7, 10},
		3,
		1,
	},
}

func Test_BinarySearch(t *testing.T) {
	ast := assert.New(t)
	for _, q := range searchQs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, BinarySearch(q.para1, q.para2), "输入：%v\n", q)
	}
}
