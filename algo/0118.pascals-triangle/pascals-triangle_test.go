package problem0118

import "testing"

var qs = []struct {
	para int
	ans  [][]int
}{
	{
		5,
		[][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		},
	},
}

func Test_Generate(t *testing.T) {

}
