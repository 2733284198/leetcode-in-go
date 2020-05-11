package problem0230

import (
	"fmt"

	"github.com/aierui/leetcode-in-go/kit"
)

type TreeNode = kit.TreeNode

func kthSmallest(root *TreeNode, k int) int {
	nums := tree2Inorder(root)

	fmt.Println(nums)

	return nums[k-1]
}

// Time O(N)
// Space O(N)
func tree2Inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	result := tree2Inorder(root.Left)
	result = append(result, root.Val)
	result = append(result, tree2Inorder(root.Right)...)

	return result
}

func kthSmallestItera(root *TreeNode, k int) int {
	var result int
	stack := make([]*TreeNode, 0)
	for k > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		k--
		root = stack[len(stack)-1]
		result = root.Val
		stack = stack[:len(stack)-1]
		root = root.Right
	}

	return result
}
