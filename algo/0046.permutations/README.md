# [46. Permutations](https://leetcode-cn.com/problems/permutations/)

## 题目

给定一个没有重复数字的序列，返回其所有可能的全排列。


示例1:

```
输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
```


## 解题思路

全排列问题 => DFS

思路一：回溯法(backTrack)

回溯法三个步骤

- choose  做选择
- explore backTrack(路径，选择列表)
- un-choose 取消选择

回溯法=递归+限制条件=暴力搜索+条件



## 总结