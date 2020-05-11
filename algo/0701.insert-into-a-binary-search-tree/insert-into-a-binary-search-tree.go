package problem0701

import "github.com/aierui/leetcode-in-go/kit"

type TreeNode = kit.TreeNode

// Time: O(H) H 指树的高度，avg O(Log(N)), worst O(N)
// insertIntoBST by recursion
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}
