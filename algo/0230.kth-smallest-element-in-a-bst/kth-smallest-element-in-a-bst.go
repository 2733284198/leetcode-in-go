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
