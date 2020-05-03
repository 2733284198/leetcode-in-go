# Go

## 调试

## 压测

golang benchmark 输出的信息说明

```
Running tool: /Users/jinrong6/develop/go-core-lib/1.14.2/bin/go test -benchmem -run=^$ github.com/aierui/leetcode-in-go/algo/0300.longest-incresing-subquence -bench ^(Benchmark_LengthOfLISByBinarySearch)$

goos: darwin
goarch: amd64
pkg: github.com/aierui/leetcode-in-go/algo/0300.longest-incresing-subquence
Benchmark_LengthOfLISByBinarySearch-4   	 6616678	       175 ns/op	     128 B/op	       2 allocs/op
PASS
ok  	github.com/aierui/leetcode-in-go/algo/0300.longest-incresing-subquence	1.620s
Success: Benchmarks passed.
```


说明


1. `6616678`  is the number of iterations  `for i := 0; i < b.N; i++ {`.
2. `175 ns/op` is approximate time it took for one iteration to complete.
3. `128 B/op` is how many bytes were allocated per op.
4. `2 allocs/op` means how many distinct memory allocations occurred per op (single iteration).

