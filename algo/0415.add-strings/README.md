# [415. Add Strings](https://leetcode-cn.com/problems/add-strings/)

## 题目

给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。

你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。



## 解题思路

注意：字符串的长度，可能存在超过 MaxInt64

步骤

- 比较较长的字符串
- 遍历较长字符串byte
- byte + byte
- 处理加法进位问题
- 特殊处理最高位进位
- []bytes -> string


## 总结