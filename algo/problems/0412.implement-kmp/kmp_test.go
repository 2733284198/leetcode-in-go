package problems0412

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 string
	para2 string
	ans   int
}{
	{
		"abcd",
		"bc",
		1,
	},
	{
		"cocacola",
		"co",
		0,
	},
	{
		"AABAACAADAABAABAAABAACAADAABAABA",
		"BA",
		2,
	},
}

var qsNext = []struct {
	para1 string
	para2 string
	ans   int
}{
	{
		"abcdbc",
		"bc",
		4,
	},
	{
		"cocacola",
		"co",
		4,
	},
	{
		"AABAACAADAABAABAAABAACAADAABAABA",
		"BA",
		30,
	},
}

func Test_SearchString(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, searchString(q.para1, q.para2), "输入：%v\n", q)
	}
}

func Benchmark_SearchString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			searchString(q.para1, q.para2)
		}
	}
}

func Test_SearchNext(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qsNext {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, searchNext(q.para1, q.para2), "输入：%v\n", q)
	}
}

func Benchmark_SearchNext(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qsNext {
			searchNext(q.para1, q.para2)
		}
	}
}
