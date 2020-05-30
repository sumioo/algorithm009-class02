/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (65.44%)
 * Likes:    84
 * Dislikes: 0
 * Total Accepted:    20.8K
 * Total Submissions: 31.8K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
 *
 * 例如，给定一个 3叉树 :
 *
 *
 *
 *
 *
 *
 *
 * 返回其层序遍历:
 *
 * [
 * ⁠    [1],
 * ⁠    [3,2,4],
 * ⁠    [5,6]
 * ]
 *
 *
 *
 *
 * 说明:
 *
 *
 * 树的深度不会超过 1000。
 * 树的节点总数不会超过 5000。
 *
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Queue []*Node

func (q *Queue) add(node *Node) {
	*q = append(*q, node)
}

func (q *Queue) remove() (*Node, error) {
	if len(*q) == 0 {
		return nil, fmt.Errorf("empty queue")
	}
	old := *q
	x := old[0]
	*q = old[1:]
	return x, nil
}

func levelOrder(root *Node) [][]int {
	q := Queue{}
	values := [][]int{}
	q.add(root)
	for len(q) > 0 {
		n := len(q)
		levelValues := []int{}
		for i := 0; i < n; i++ {
			node, _ := q.remove()
			if node == nil {
				continue
			}
			levelValues = append(levelValues, node.Val)
			for _, child := range node.Children {
				q.add(child)
			}
		}
		values = append(values, levelValues)
	}
	return values
}

// @lc code=end

