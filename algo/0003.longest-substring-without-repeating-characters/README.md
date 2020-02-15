# [3. Longest-Substring-Without-Repeating-Characters](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)

## 题目

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例1:

```
输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```

示例2:

```
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
```


示例 3:

```
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
```


## 解题思路

利用 `s[left:i+1]` 来表示 `s[:i+1]` 中的包含 `s[i]` 的最长子字符串。 `location[s[i]]` 是字符 `s[i]` 在 `s[:i+1]` 中倒数第二次出现的序列号。 当 `left < location[s[i]]` 的时候，说明字符 `s[i]` 出现了两次。需要设置 `left = location[s[i]] + 1`, 保证字符 `s[i]` 只出现一次。


## 总结

此题为经典的窗口滑动中非定长算法。

解题另一个思路为抽象符合条件的字符串在窗口内，依次遍历字符串，当前位置字符串如果包含在窗口中，窗口中的开始位置更新为重复位置的下一位 最后得到最长非重复字符串的值。

```go
strings.Index(s[j:i], string(s[i]))
```


但是 golang 的 byte 操作性能远远高于操作string的性能



