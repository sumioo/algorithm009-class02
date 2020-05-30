/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Medium (65.52%)
 * Likes:    264
 * Dislikes: 0
 * Total Accepted:    108.9K
 * Total Submissions: 166K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 前序 遍历。
 *
 * 示例:
 *
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * 输出: [1,2,3]
 *
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
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

func helper(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	*values = append(*values, root.Val)
	helper(root.Left, values)
	helper(root.Right, values)
}
func preorderTraversalA(root *TreeNode) []int {
	values := []int{}
	helper(root, &values)
	return values
}

type Stack []*TreeNode

func (s *Stack) push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) pop() *TreeNode {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

func preorderTraversal(root *TreeNode) []int {
	values := []int{}
	stack := Stack{}
	currentNode := root

	for currentNode != nil || len(stack) != 0 {
		for currentNode != nil {
			values = append(values, currentNode.Val)
			stack.push(currentNode)
			currentNode = currentNode.Left
		}

		top := stack.pop()
		currentNode = top.Right
	}

	return values

}

// @lc code=end

