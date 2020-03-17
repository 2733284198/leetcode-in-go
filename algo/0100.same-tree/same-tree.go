package problem0100

import "github.com/aierui/leetcode-in-go/kit"

// TreeNode Definition for a binary tree node.
type TreeNode = kit.TreeNode

// isSameTree by recursion
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}

	if p == nil && q == nil {
		return true
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
