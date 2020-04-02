# [237. Delete Node In A Linked List](https://leetcode-cn.com/problems/delete-node-in-a-linked-list/)

## 题目

给定一个数组，包含从 1 到 N 所有的整数，但其中缺了两个数字。你能在 O(N) 时间内只用 O(1) 的空间找到它们吗？

以任意顺序返回这两个数字均可。

示例 1:

```
输入: [1]
输出: [2,3]
```

示例 2:

```
输入: [2,3]
输出: [1,4]
提示：

nums.length <= 30000
```



## 解题思路


思路一：累加做差

1. 求 [1, n] 的和记为S0， nums 数组的和记为S1
2. 那么缺少的两个数的和 a + b = S0 - S1
3. 两个数首先是不相等，mid = (a + b) / 2, 若 a < b , 则 a < mid, b > mid
4. 先找a，sum([1, mid]) - sum(nums[i] < mid)，为S1中小于 mid 的数



## 总结