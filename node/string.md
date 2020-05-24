# 字符串


## 字符串基础知识

- 字符串实现的原理（编码）？
- 字符串内存地址是否连续？
- 字符串的基本操作？
    - 长度
    - 截取
    - 搜索
    - 分割
    - 合并
    - …
- 定长字符串与变长字符串？

string immutable?

- [Are your strings immutable?](https://lemire.me/blog/2017/07/07/are-your-strings-immutable/)


## 常见字符串题目

- 最长公共子串、最长公共子序列
- 最长回文子串、最长回文子序列
- …

## 字符串匹配算法

1. Brute-Force
2. Dynamic Programming
3. Rabin-Karp Search
4. KMP
5. Boyer-Moore
6. Sunday
7. …

### Brute-Force

```go
for i:=0; i < n; i ++ {
    for j:=0; j < m; j ++ {
        if s1[i] == s2[j] {
            // todo
        } else {
            // todo
        }
    }
}
```

### Dynamic Programming

```go
dp := make([][]int, len(s1)+1)
for i := 0; i <= len(s1); i++ {
    dp[i] = make([]int, len(s2)+1)
}
```

### Rabin-Karp Search

原理：

Michael O. Rabin 和 Richard M. Karp 在1987年提出一个想法，即可以对模式串进行哈希运算并将其哈希值与文本中子串的哈希值进行比对。总的来说这一想法非常浅显，唯一的问题在于我们需要找到一个哈希函数 ，它需要能够对不同的字符串返回不同的哈希值。例如，该哈希函数可能会对每个字符的ASCII码进行算，但同时我们也需要仔细考虑对多语种文本的支持。

实现：

```go
// from go source src/strings/strings.go
// primeRK is the prime base used in Rabin-Karp algorithm.
const primeRK = 16777619

// hashStr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func hashStr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

func indexRabinKarp(s, substr string) int {
	// Rabin-Karp search
	hashss, pow := hashStr(substr)
	n := len(substr)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(s[i])
	}
	// 这里为啥需要拆开呢 ？
	// 可以直接 从 0 到 len(s) 循环遍历吗？
	if h == hashss && s[:n] == substr {
		return 0
	}
	for i := n; i < len(s); {
		h *= primeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashss && s[i-n:i] == substr {
			return i - n
		}
	}
	return -1
}
```


### KMP


原理：

生成待匹配字符串的部分匹配值（前缀表） O(n)，"前缀"指除了最后一个字符以外，一个字符串的全部头部组合；"后缀"指除了第一个字符以外，一个字符串的全部尾部组合。部分匹配值"就是"前缀"和"后缀"的最长的共有元素的长度。


实现：


参考

- http://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html

### Boyer-Moore




参考

- https://www.ruanyifeng.com/blog/2013/05/boyer-moore_string_search_algorithm.html

### Sunday