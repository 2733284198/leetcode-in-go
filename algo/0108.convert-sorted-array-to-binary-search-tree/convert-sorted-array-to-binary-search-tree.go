package problem0108

import (
	"github.com/aierui/leetcode-in-go/kit"
)

type TreeNode = kit.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {

	if nums == nil || len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	left := nums[:mid]
	right := nums[mid+1:]

	node := &TreeNode{
		nums[mid],
		sortedArrayToBST(left),
		sortedArrayToBST(right),
	}

	return node
}
