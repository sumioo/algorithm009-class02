/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (62.88%)
 * Likes:    529
 * Dislikes: 0
 * Total Accepted:    144.7K
 * Total Submissions: 230K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
 *
 *
 *
 * 示例：
 * 二叉树：[3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回其层次遍历结果：
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
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
# 深度优先遍历
我们可以将slice索引和节点层次联系起来，比如节点层次是0,那么我们将节点值填入索引为0的格子。
有一个问题是我们事先并不知道树有多少层，无法得知我们应该创建多长的slice。显然我们可以通过深度优先搜索得到树的深度
然后创建相应长度的slice，这样就不用担心slice越界了。进一步思考我们可以知道，深度遍历时时一层一层往下探的，
绝对不会跳跃，那么我们可以利用这一点，但我们发现slice长度等于level时（以0作为level的起点），我们用append增加slice长度。

*/

func dfs(root *TreeNode, level int, values *[][]int) {
	if root == nil {
		return
	}

	if len(*values) == level {
		*values = append(*values, []int{root.Val})
	} else {
		(*values)[level] = append((*values)[level], root.Val)
	}
	dfs(root.Left, level+1, values)
	dfs(root.Right, level+1, values)
}

func levelOrderDFS(root *TreeNode) [][]int {
	values := [][]int{}
	dfs(root, 0, &values)
	return values
}

/*
# 广度优先遍历


*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	values := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		levelValues := []int{}
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			levelValues = append(levelValues, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		values = append(values, levelValues)
	}
	return values
}

// @lc code=end

