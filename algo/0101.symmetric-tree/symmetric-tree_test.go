package problem0101

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 []int
	ans   bool
}{
	{
		[]int{},
		[]int{},
		true,
	},
	{
		[]int{1, 2, 2},
		[]int{2, 1, 2},
		true,
	},

	{
		[]int{1, 2, 3, 2, 3},
		[]int{2, 3, 1, 2, 3},
		false,
	},
}

func Test_isSymmetric(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, isSymmetric(kit.PreIn2Tree(q.para1, q.para2)), "输入：%v\n", q)
	}
}

func Test_isSymmetricByBFS(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, isSymmetricByBFS(kit.PreIn2Tree(q.para1, q.para2)), "输入：%v\n", q)
	}
}

func Test_isSymmetricByChan(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, isSymmetricByChan(kit.PreIn2Tree(q.para1, q.para2)), "输入：%v\n", q)
	}
}
