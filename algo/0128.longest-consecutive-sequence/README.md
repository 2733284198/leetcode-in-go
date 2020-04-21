# [128. Longest Consecutive Sequence](https://leetcode-cn.com/problems/longest-consecutive-sequence/)

## 题目

给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

```
输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
```



## 解题思路

时间复杂度：O(n)；遍历n个元素录入map，再实际在查找最长length的过程中再遍历一遍
空间复杂度：O(n)；将所有数组中n个元素录入map


## 总结