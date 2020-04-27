package problem0102

import "github.com/aierui/leetcode-in-go/kit"

type TreeNode = kit.TreeNode

var result [][]int

func levelOrder(root *TreeNode) [][]int {
	result = [][]int{}
	dfs(root, 0)

	return result
}

func dfs(root *TreeNode, level int) {
	if root != nil {
		if len(result) == level {
			result = append(result, []int{})
		}
		result[level] = append(result[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
}
