package problems0421

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs62 = []struct {
	para string
	ans  int
}{
	{
		"A40",
		138632,
	},
	{
		"FX8",
		161270,
	},
	{
		"kAKOpmZUD",
		4496239229889483,
	},
	{
		"kALgDSVP8",
		4496265096088262,
	},
}

var qs10 = []struct {
	para int
	ans  string
}{
	{
		138632,
		"A40",
	},
	{
		161270,
		"FX8",
	},
	{
		4496239229889483,
		"kAKOpmZUD",
	},
	{
		4496265096088262,
		"kALgDSVP8",
	},
}

func Test_Base62To10(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs62 {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, base62To10(q.para), "输入：%v\n", q)
	}
}

func Benchmark_Base62To10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs62 {
			base62To10(q.para)
		}
	}
}

func Test_Base10To62(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs10 {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, base10To62(q.para), "输入：%v\n", q)
	}
}

func Benchmark_Base10To62(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs10 {
			base10To62(q.para)
		}
	}
}
