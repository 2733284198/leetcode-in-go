package problem0094

import "github.com/aierui/leetcode-in-go/kit"

type TreeNode = kit.TreeNode

var recursionRtn []int

func inOrderTraversalByRecursion(root *TreeNode) []int {
	recursionRtn = make([]int, 0)
	recursion(root)

	return recursionRtn
}

func recursion(root *TreeNode) {
	if root != nil {
		recursion(root.Left)
		recursionRtn = append(recursionRtn, root.Val)
		recursion(root.Right)
	}
}
