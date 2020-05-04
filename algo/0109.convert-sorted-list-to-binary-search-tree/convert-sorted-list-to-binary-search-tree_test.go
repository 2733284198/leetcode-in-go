package problem0109

import (
	"fmt"
	"testing"
)

var qs = []struct {
	para []int
	ans  []int
}{
	{
		[]int{-10, -3, 0, 5, 9},
		[]int{-10, -3, 0, 5, 9},
	},
}

func Test_SortedListToBST(t *testing.T) {
	//ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		// list := kit.Ints2List(q.para)
		//ast.Equal(q.ans, kit.Tree2Inorder(sortedListToBST(kit.Ints2List(q.para))), "输入：%v\n", q)
	}
}
