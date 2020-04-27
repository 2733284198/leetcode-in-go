package problem0107

import "github.com/aierui/leetcode-in-go/kit"

type TreeNode = kit.TreeNode

var result [][]int

func levelOrderBottom(root *TreeNode) [][]int {
	result = [][]int{}
	dfs(root, 0)

	// reverse array
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

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
