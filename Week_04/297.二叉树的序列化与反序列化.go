/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 *
 * https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/description/
 *
 * algorithms
 * Hard (47.23%)
 * Likes:    202
 * Dislikes: 0
 * Total Accepted:    24.5K
 * Total Submissions: 51.6K
 * Testcase Example:  '[1,2,3,null,null,4,5]'
 *
 *
 * 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
 *
 * 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 /
 * 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
 *
 * 示例:
 *
 * 你可以将以下二叉树：
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * ⁠    / \
 * ⁠   4   5
 *
 * 序列化为 "[1,2,3,null,null,4,5]"
 *
 * 提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode
 * 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
 *
 * 说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。
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

import "strconv"
import "strings"

type Codec struct {
	emptyNode string
}

func Constructor() Codec {
	return Codec{emptyNode: "None"}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var preorder func(root *TreeNode, serializedNodes *[]string)
	serializedNodes := []string{}

	preorder = func(root *TreeNode, serializedNodes *[]string) {
		if root == nil {
			*serializedNodes = append(*serializedNodes, this.emptyNode)
			return
		}
		*serializedNodes = append(*serializedNodes, strconv.Itoa(root.Val))
		preorder(root.Left, serializedNodes)
		preorder(root.Right, serializedNodes)
	}
	preorder(root, &serializedNodes)
	return strings.Join(serializedNodes, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	serializedNodes := strings.Split(data, ",")

	var rdeserialize func(serializedNodes *[]string) *TreeNode

	rdeserialize = func(serializedNodes *[]string) *TreeNode {
		if (*serializedNodes)[0] == this.emptyNode {
			*serializedNodes = (*serializedNodes)[1:]
			return nil
		}
		val, _ := strconv.Atoi((*serializedNodes)[0])
		root := &TreeNode{Val: val}
		*serializedNodes = (*serializedNodes)[1:]
		root.Left = rdeserialize(serializedNodes)
		root.Right = rdeserialize(serializedNodes)
		return root
	}
	return rdeserialize(&serializedNodes)
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end

