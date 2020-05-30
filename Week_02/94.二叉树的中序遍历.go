/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (71.43%)
 * Likes:    511
 * Dislikes: 0
 * Total Accepted:    162.7K
 * Total Submissions: 227.5K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的中序 遍历。
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
 * 输出: [1,3,2]
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

//左子树-根-右子树

func helper(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, values)
	*values = append(*values, root.Val)
	helper(root.Right, values)
}

func inorderTraversalA(root *TreeNode) []int {
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

func (s Stack) top() *TreeNode {
	return s[len(s)-1]
}

func inorderTraversal(root *TreeNode) []int {
	values := []int{}
	stack := Stack{}
	currentNode := root
	for currentNode != nil || len(stack) != 0 {
		for currentNode != nil {
			stack.push(currentNode)
			currentNode = currentNode.Left
		}
		top := stack.pop()
		values = append(values, top.Val)
		currentNode = top.Right
	}
	return values
}

// @lc code=end

