/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (31.16%)
 * Likes:    598
 * Dislikes: 0
 * Total Accepted:    125.3K
 * Total Submissions: 400K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 *
 * 假设一个二叉搜索树具有如下特征：
 *
 *
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 *
 *
 * 示例 1:
 *
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 * 根节点的值为 5 ，但是其右子节点值为 4 。
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

/*
二叉树的定义
根节点大于左子树节点，根节点大于右子树节点，且左右子树根节点具有相同的性质
*/

import "math"

func validBST(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return validBST(root.Left, lower, root.Val) && validBST(root.Right, root.Val, upper)
}

func isValidBSTA(root *TreeNode) bool {
	return validBST(root, math.MinInt64, math.MaxInt64)
}

/*
搜索二叉树的中序遍历是有序递增的，因此可以判断树是否搜索二叉树，自然地我们会想到使用中序遍历和一个数组来保存遍历结果，再判断数组是否递增的。
但是想到我们判断递增只需用当前树和前一个数相比较即可，如果发现不符合要求，立即返回即可。所以可以维护一个前一个值得变量，直接在中序遍历时比较即可。
*/
func isValidBST(root *TreeNode) bool {
	var lastVal int = math.MinInt64
	var inOrder func(root *TreeNode) bool
	inOrder = func(root *TreeNode) bool {
		if root == nil {
			return true
		}

		if !inOrder(root.Left) {
			return false
		}
		if lastVal >= root.Val {
			return false
		}
		lastVal = root.Val
		if !inOrder(root.Right) {
			return false
		}
		return true
	}

	return inOrder(root)
}

// @lc code=end


