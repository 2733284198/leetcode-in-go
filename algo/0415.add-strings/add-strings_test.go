package problem0415

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 string
	para2 string
	ans   string
}{
	{
		"12345",
		"54321",
		"66666",
	},
}

func Test_AddStrings(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
	}
}
