package problem0099

import "github.com/aierui/leetcode-in-go/kit"

// TreeNode
type TreeNode = kit.TreeNode

var last, first, second *TreeNode

// recoverTree 中序递归
func recoverTree(root *TreeNode) {
	last, first, second = nil, nil, nil
	dfs(root)
	first.Val, second.Val = second.Val, first.Val
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if last != nil && root.Val <= last.Val {
		if first != nil {
			second = root
			return //剪枝
		}
		first, second = last, root
	}
	last = root

	dfs(root.Right)
}
