package problem0110

import "github.com/aierui/leetcode-in-go/kit"

// TreeNode
type TreeNode = kit.TreeNode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBalanced(root.Left) && isBalanced(root.Right) && abs(height(root.Left), height(root.Right)) <= 1
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(height(node.Left), height(node.Right)) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
