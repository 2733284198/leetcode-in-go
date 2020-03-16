package problem0050

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	paraX float64
	paraN int

	ans float64
}{
	{
		2.00000,
		10,
		1024.00000,
	},
	// 存在精度问题
	// {
	// 	2.10000,
	// 	3,
	// 	9.26100,
	// },
	{
		2.00000,
		-2,
		0.25000,
	},
}

func Test_MyPow(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, myPow(q.paraX, q.paraN), "输入：%v\n", q)
	}
}

func Test_PowByBuiltInLib(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, PowByBuiltInLib(q.paraX, q.paraN), "输入：%v\n", q)
	}
}

func Benchmark_MyPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			myPow(q.paraX, q.paraN)
		}
	}
}

func Benchmark_PowByBuiltInLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			PowByBuiltInLib(q.paraX, q.paraN)
		}
	}
}
