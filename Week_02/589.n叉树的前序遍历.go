/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/
 *
 * algorithms
 * Easy (73.05%)
 * Likes:    78
 * Dislikes: 0
 * Total Accepted:    27.4K
 * Total Submissions: 37.5K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的前序遍历。
 *
 * 例如，给定一个 3叉树 :
 *
 *
 *
 *
 *
 *
 *
 * 返回其前序遍历: [1,3,5,6,2,4]。
 *
 *
 *
 * 说明: 递归法很简单，你可以使用迭代法完成此题吗?
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func helper(root *Node, values *[]int) {
	if root == nil {
		return
	}
	*values = append(*values, root.Val)
	for _, child := range root.Children {
		helper(child, values)
	}
}

func preorderA(root *Node) []int {
	values := &[]int{}
	helper(root, values)
	return *values
}

type Stack []*Node

func (s *Stack) push(node *Node) {
	*s = append(*s, node)
}

func (s *Stack) pop() *Node {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

func preorder(root *Node) []int {
	stack := Stack{}
	stack.push(root)
	values := []int{}
	for len(stack) > 0 {
		node := stack.pop()
		if node == nil {
			continue
		}
		values = append(values, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.push(node.Children[i])
		}
	}
	return values
}

// @lc code=end

