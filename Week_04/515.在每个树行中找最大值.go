/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
 *
 * https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/description/
 *
 * algorithms
 * Medium (60.23%)
 * Likes:    67
 * Dislikes: 0
 * Total Accepted:    12.5K
 * Total Submissions: 20.7K
 * Testcase Example:  '[1,3,2,5,3,null,9]'
 *
 * 您需要在二叉树的每一行中找到最大的值。
 *
 * 示例：
 *
 *
 * 输入:
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       3   2
 * ⁠      / \   \
 * ⁠     5   3   9
 *
 * 输出: [1, 3, 9]
 *
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

func dfs(root *TreeNode, level int, values *[]int) {
	if root == nil {
		return
	}

	if len(*values) == level {
		*values = append(*values, root.Val)
	} else {
		if (*values)[level] < root.Val {
			(*values)[level] = root.Val
		}
	}
	dfs(root.Left, level+1, values)
	dfs(root.Right, level+1, values)
}

func largestValuesDFS(root *TreeNode) []int {
	values := []int{}
	dfs(root, 0, &values)
	return values
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	values := []int{}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		max := queue[0].Val
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if max < node.Val {
				max = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		values = append(values, max)
	}
	return values
}

// @lc code=end

