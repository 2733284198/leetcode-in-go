# [98. Validate Binary Search Tree](https://leetcode-cn.com/problems/validate-binary-search-tree/)

## 题目

给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

示例 1:

```
输入:
    2
   / \
  1   3
输出: true
```

示例 2:

```
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
```

解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。



## 解题思路


思路一：前序递归
    1. 递归方向 根-左-右
    2. 上下限 min max 计算

思路二：前序迭代（stack）
    和 前序递归逻辑类似，需要自己维护 stack 来控制 root min max 出入栈

思路三：中序递归
    1. 递归方向 左-根-右
    2. 上一节点 last 赋值时机
    3. last.Val 必须小于 root.Val 

## 总结