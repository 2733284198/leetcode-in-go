package problem0095

import "github.com/aierui/leetcode-in-go/kit"

type TreeNode = kit.TreeNode

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return doGenerateTrees(1, n)
}

func doGenerateTrees(left, right int) []*TreeNode {
	if left > right {
		return []*TreeNode{nil}
	}

	if left == right {
		return []*TreeNode{&TreeNode{Val: left}}
	}

	trees := make([]*TreeNode, 0, right-left+1)

	for i := left; i <= right; i++ {
		leftTrees := doGenerateTrees(left, i-1)
		rightTrees := doGenerateTrees(i+1, right)

		for m := 0; m < len(leftTrees); m++ {
			for n := 0; n < len(rightTrees); n++ {
				tree := &TreeNode{Val: i}

				tree.Left = leftTrees[m]
				tree.Right = rightTrees[n]

				trees = append(trees, tree)
			}
		}
	}

	return trees
}
