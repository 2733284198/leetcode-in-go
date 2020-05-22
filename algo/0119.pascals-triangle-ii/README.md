# [119. Pascals Triangle ii](https://leetcode-cn.com/problems/pascals-triangle-ii/)

## 题目

给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。

![](https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif)

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

```
输入: 3
输出: [1,3,3,1]
```
你可以优化你的算法到 O(k) 空间复杂度吗？

## 解题思路

思路：只利用一个[]int，行数下移即在尾部添加元素1，然后倒序计算当前行


## 总结
