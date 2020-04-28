package problem0111

import "github.com/aierui/leetcode-in-go/kit"

type TreeNode = kit.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return height(root)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return min(height(node.Left), height(node.Right)) + 1
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
