package problem0071

import "github.com/stretchr/testify/assert"

var qs = []struct {
	para string
	ans  string
}{ 
	{
		"/home/",
		"/home",
	},
	{
		"/../",
		"/"
	},
	{
		"/home//foo",
		"/home/foo"
	},
	{
		"/a/./b/../../c/",
		"/c"
	},
	{
		"/a/../../b/../c//.//",
		"/c",
	},
	{
		"/a//b////c/d//././/..",
		"/a/b/c",
	},
}


func Test_SimplifyPath(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n",q)
		ast.Equal(q.ans, simplifyPath(q.para), "输入：%v\n", q)
	}
}
