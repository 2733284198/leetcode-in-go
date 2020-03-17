package problem0101

import "github.com/aierui/leetcode-in-go/kit"

type TreeNode = kit.TreeNode

// isSymmetric by recursion
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	} else if root.Left == nil && root.Right == nil {
		return true
	} else if root.Left == nil || root.Right == nil {
		return false
	}

	if root.Left.Val != root.Right.Val {
		return false
	}

	tmp := root.Right.Right
	root.Right.Right = root.Left.Right
	root.Left.Right = tmp

	return isSymmetric(root.Left) && isSymmetric(root.Right)
}

func isSymmetricByBFS(root *TreeNode) bool {
	lq := []*TreeNode{}
	rq := []*TreeNode{}

	lq = append(lq, root)
	rq = append(rq, root)

	for len(lq) != 0 && len(rq) != 0 {
		lcur, rcur := lq[0], rq[0]
		lq, rq = lq[1:], rq[1:]

		if lcur == nil && rcur == nil {
			continue
		} else if lcur != nil && rcur != nil && lcur.Val == rcur.Val {
			lq = append(lq, lcur.Left, lcur.Right)
			rq = append(rq, rcur.Right, rcur.Left)
		} else {
			return false
		}
	}

	if len(lq) == 0 && len(rq) == 0 {
		return true
	}

	return false
}

// isSymmetricByChan
// 思路是开两个协程分别生产两种递归顺序，然后主线程依次接收，并
// 对比每次产生的顺序是否一致，如果始终一致，表示是对称二叉树。
func isSymmetricByChan(root *TreeNode) bool { // 使用通道 chan 构造迭代器，判断每次结果是否一样
	nodeCount := getNodeCount(root) // 总节点数（加上nil节点）
	lchan := make(chan int)
	rchan := make(chan int)
	go leftGen(lchan, root) // 开协程生成遍历顺序
	go rightGen(rchan, root)
	for nodeCount > 0 {
		lcur := <-lchan // 从通道中分别取出遍历顺序
		rcur := <-rchan
		if lcur != rcur { // 判断遍历顺序是否相同
			return false
		}
		nodeCount--
	}
	defer close(lchan)
	defer close(rchan)
	return true
}

func getNodeCount(cur *TreeNode) int {
	if cur == nil {
		return 1
	}
	return getNodeCount(cur.Left) + getNodeCount(cur.Right) + 1
}

func leftGen(lchan chan int, cur *TreeNode) {
	if cur == nil {
		lchan <- -1
		return
	}
	lchan <- cur.Val // 加入左遍历顺序通道
	leftGen(lchan, cur.Left)
	leftGen(lchan, cur.Right)
}

func rightGen(rchan chan int, cur *TreeNode) {
	if cur == nil {
		rchan <- -1
		return
	}
	rchan <- cur.Val // 加入右遍历顺序通道
	rightGen(rchan, cur.Right)
	rightGen(rchan, cur.Left)
}
