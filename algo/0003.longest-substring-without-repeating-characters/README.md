# [2. Longest-Substring-Without-Repeating-Characters](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)

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
