package problem0144

import "github.com/aierui/leetcode-in-go/kit"

type TreeNode = kit.TreeNode

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := []int{root.Val}

	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)

	return res
}

func preorderTraversalByStack(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	slice := make([]int, 0)
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		slice = append(slice, p.Val)

		if p.Right != nil {
			stack = append(stack, p.Right)
		}

		if p.Left != nil {
			stack = append(stack, p.Left)
		}

	}

	return slice
}
