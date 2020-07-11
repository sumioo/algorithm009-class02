/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (47.58%)
 * Likes:    742
 * Dislikes: 0
 * Total Accepted:    78.8K
 * Total Submissions: 158.2K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 *
 * 获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) -
 * 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
 *
 *
 *
 * 进阶:
 *
 * 你是否可以在 O(1) 时间复杂度内完成这两种操作？
 *
 *
 *
 * 示例:
 *
 * LRUCache cache = new LRUCache( 2 缓存容量 );
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回  1
 * cache.put(3, 3);    // 该操作会使得关键字 2 作废
 * cache.get(2);       // 返回 -1 (未找到)
 * cache.put(4, 4);    // 该操作会使得关键字 1 作废
 * cache.get(1);       // 返回 -1 (未找到)
 * cache.get(3);       // 返回  3
 * cache.get(4);       // 返回  4
 *
 *

*/

package week8

// @lc code=start

type ListNode struct {
	key   int
	value int
	prev  *ListNode
	next  *ListNode
}

type List struct {
	head *ListNode
	tail *ListNode
}

func (l *List) insertBefore(new *ListNode, node *ListNode) {
	new.prev = node.prev
	new.next = node
	node.prev.next = new
	node.prev = new
}

func (this *List) moveToHead(node *ListNode) {
	this.remove(node)
	this.insertBefore(node, this.head.next)
}

func (this *List) insertFirst(key int, value int) *ListNode {
	new := &ListNode{key: key, value: value}
	after := this.head.next
	this.insertBefore(new, after)
	return new
}

func (this *List) moveToTail(node *ListNode) {
	this.remove(node)
	this.insertBefore(node, this.tail)
}

func (this *List) remove(node *ListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
}

func (this *List) removeTail() {
	if this.tail.prev == this.head {
		return
	}
	this.remove(this.tail.prev)
}

func (this *List) getTail() *ListNode {
	if this.tail.prev == this.head {
		return nil
	}
	return this.tail.prev
}

type LRUCache struct {
	capacity int
	cache    map[int]*ListNode
	list     *List
	size     int
}

func Constructor(capacity int) LRUCache {
	dummyHead := &ListNode{}
	dummyTail := &ListNode{}
	dummyHead.next = dummyTail
	dummyTail.prev = dummyHead
	list := &List{head: dummyHead, tail: dummyTail}
	return LRUCache{capacity: capacity, cache: make(map[int]*ListNode, capacity), list: list}
}

//Get 查询并更新
func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.list.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		node.value = value
		this.list.moveToHead(node)
		return
	}

	if this.size == this.capacity {
		tail := this.list.getTail()
		this.list.removeTail()
		delete(this.cache, tail.key)
		this.size--
	}
	new := this.list.insertFirst(key, value)
	this.cache[key] = new
	this.size++

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
