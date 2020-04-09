# [65. Valid Number](https://leetcode-cn.com/problems/valid-number/)

## 题目

验证给定的字符串是否可以解释为十进制数字。

例如:

```
"0" => true
" 0.1 " => true
"abc" => false
"1 a" => false
"2e10" => true
" -90e3   " => true
" 1e" => false
"e3" => false
" 6e-1" => true
" 99e2.5 " => false
"53.5e93" => true
" --6 " => false
"-+3" => false
"95a54e53" => false
```

说明: 我们有意将问题陈述地比较模糊。在实现代码之前，你应当事先思考所有可能的情况。这里给出一份可能存在于有效十进制数字中的字符列表：

- 数字 0-9
- 指数 - "e"
- 正/负号 - "+"/"-"
- 小数点 - "."

当然，在输入中，这些字符的上下文也很重要。



## 解题思路


思路一：有限状态机

![](https://pic.leetcode-cn.com/e87dc057b4d26e4b8554b4fb311ffa9df8301ab5f10d3755a1ab115d8cf6c925-image.png)


[有限状态机](https://zh.wikipedia.org/zh-hans/%E6%9C%89%E9%99%90%E7%8A%B6%E6%80%81%E6%9C%BA)




## 总结