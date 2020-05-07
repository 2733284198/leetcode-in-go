package problem0145

import "github.com/aierui/leetcode-in-go/kit"

type TreeNode = kit.TreeNode

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := []int{}

	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)

	return res
}

func postorderTraversalByStack(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	slice := make([]int, 0)
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		root := stack[len(stack)-1]

		if root.Right != nil {
			stack = append(stack, root.Right)
		}

		if root.Left != nil {
			stack = append(stack, root.Left)
		}

		if root.Left == nil && root.Right == nil {
			stack = stack[:len(stack)-1]
			slice = append(slice, root.Val)
		}

		root.Left = nil
		root.Right = nil
	}

	return slice
}
