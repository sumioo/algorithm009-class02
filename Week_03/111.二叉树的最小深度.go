/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (42.37%)
 * Likes:    261
 * Dislikes: 0
 * Total Accepted:    79.2K
 * Total Submissions: 186.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 *
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最小深度  2.
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDepthA(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var mindepth int
	if root.Left == nil && root.Right != nil {
		mindepth = minDepth(root.Right) + 1
	} else if root.Left != nil && root.Right == nil {
		mindepth = minDepth(root.Left) + 1
	} else {
		mindepth = min(minDepth(root.Left), minDepth(root.Right)) + 1
	}

	return mindepth
}

type Wrapper struct {
	node  *TreeNode
	level int
}

type Queue []*Wrapper

func (q *Queue) add(w *Wrapper) {
	*q = append(*q, w)
}

func (q *Queue) remove() (*Wrapper, error) {
	old := *q
	n := len(old)
	if n == 0 {
		return nil, fmt.Errorf("empty queue")
	}
	w := old[0]
	*q = old[1:]
	return w, nil
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := Queue{}
	q.add(&Wrapper{node: root, level: 1})

	for len(q) > 0 {
		n := len(q)

		for i := 0; i < n; i++ {
			w, _ := q.remove()
			if w.node.Left == nil && w.node.Right == nil {
				return w.level
			}

			if w.node.Left != nil {
				q.add(&Wrapper{node: w.node.Left, level: w.level + 1})
			}
			if w.node.Right != nil {
				q.add(&Wrapper{node: w.node.Right, level: w.level + 1})
			}
		}
	}
	return 0
}

// @lc code=end

