/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	l3 := &ListNode{}
	head := l3

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l3.Next = l1
			l1 = l1.Next
		} else {
			l3.Next = l2
			l2 = l2.Next
		}

		l3 = l3.Next
	}

	if l1 == nil {
		l3.Next = l2
	} else {
		l3.Next = l1
	}

	return head.Next

}

// @lc code=end

// @lc code=start

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
  if l1 == nil {
    return l2
  } else if l2 == nil {
    return l1
  } else if l1.Val < l2.Val {
    return l1.Next = mergeTwoLists(l1.Next, l2)
  } else {
    return l2.Next = mergeTwoLists(l2.Next, l1)
  }

}
// @lc code=end
