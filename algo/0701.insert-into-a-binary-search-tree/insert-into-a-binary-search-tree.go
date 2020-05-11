package problem0701

import "github.com/aierui/leetcode-in-go/kit"

type TreeNode = kit.TreeNode

// Time: O(H) H 指树的高度，avg O(Log(N)), worst O(N)
// Space: avg O(H) worst O(N)
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

func insertBSTItera(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	r := root

	for {
		if r.Val > val {
			if r.Left == nil {
				r.Left = &TreeNode{Val: val}
				break
			} else {
				r = r.Left
			}
		} else {
			if r.Right == nil {
				r.Right = &TreeNode{Val: val}
				break
			} else {
				r = r.Right
			}
		}
	}

	return root
}
