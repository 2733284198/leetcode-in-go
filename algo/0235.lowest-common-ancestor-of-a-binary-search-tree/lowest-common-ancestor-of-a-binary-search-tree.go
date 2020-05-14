package problem0235

import "github.com/aierui/leetcode-in-go/kit"

type TreeNode = kit.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pVal, qVal := p.Val, q.Val
	node := root
	for node != nil {
		parentVal := node.Val

		switch {
		case pVal > parentVal && qVal > parentVal:
			node = node.Right
		case pVal < parentVal && qVal < parentVal:
			node = node.Left
		default:
			return node
		}
	}
	return nil
}
