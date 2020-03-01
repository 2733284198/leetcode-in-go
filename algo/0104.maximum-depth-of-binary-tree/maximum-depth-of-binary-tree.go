package problem0104

import "github.com/aierui/leetcode-in-go/kit"

// TreeNode for a binary tree node.
type TreeNode = kit.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lH, rH := 0, 0

	if root.Left != nil {
		lH = maxDepth(root.Left)
	}

	if root.Right != nil {
		rH = maxDepth(root.Right)
	}

	var max int
	if lH > rH {
		max = lH
	} else {
		max = rH
	}

	return max + 1
}
