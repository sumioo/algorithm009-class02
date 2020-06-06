/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (65.50%)
 * Likes:    516
 * Dislikes: 0
 * Total Accepted:    84.1K
 * Total Submissions: 125.4K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
前序遍历 根节点|左子树|右子树
中序遍历 左子树|根节点|右子树

要构造一个节点，那么我们必须先知道节点的值，然后是它的左子树和右子树
前序遍历总是先访问根节点，然后是左子树，接着是右子树，同样的左子树和右子树第一个元素也是当前子树的根节点。所以我们可以得知前序遍历的preorder[0]
就是根节点，但从前序遍历我们不能区分左子树|右子树的分界点，再看看中序遍历，中序遍历总是先访问左子树|根节点|右子树，也就是说根节点
把中序遍历根数组分为了左子树部分和右子树部分，从而我们可以知道左子树和右子树的长度，知道长度后我们就可以划分前序遍历的左子树和右子树了

*/

func index(nums []int, num int) int {
	for index, n := range nums {
		if n == num {
			return index
		}
	}
	return -1
}

func helper(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	i := index(inorder, rootVal)
	root.Left = helper(preorder[1:i+1], inorder[:i])
	root.Right = helper(preorder[i+1:], inorder[i+1:])
	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return helper(preorder, inorder)
}

// @lc code=end

