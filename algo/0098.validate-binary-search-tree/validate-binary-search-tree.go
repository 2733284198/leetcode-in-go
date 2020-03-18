package problem0098

import (
	"math"

	"github.com/aierui/leetcode-in-go/kit"
)

type TreeNode = kit.TreeNode
// isValidBST by recursion
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(root *TreeNode, left, right int) bool  {
	
	if root == nil {
		return true
	}

	if left >= root.Val || right <= root.Val {
		return false
	}

	return isBST(root.Left, left, root.Val) && isBST(root.Right, root.Val, right)
}