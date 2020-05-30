package week2

import "strconv"

//Node tree node
type Node struct {
	key   int
	value int
	left  *Node
	right *Node
}

func (node *Node) String() string {
	return strconv.Itoa(node.key)
}

//Bst binary tree
type Bst struct {
	root *Node
}

func (b *Bst) insert(key int, value int) {
	b.root = insert(b.root, key, value)
}

func insert(root *Node, key int, value int) *Node {
	if root == nil {
		return &Node{key: key, value: value}
	}

	if root.key > key {
		root.left = insert(root.left, key, value)
	} else if root.key < key {
		root.right = insert(root.right, key, value)
	} else {
		root.value = value
	}
	return root
}

func min(root *Node) *Node {
	if root.left == nil {
		return root
	}
	return min(root.left)
}

func (b *Bst) delete(key int) {
	b.root = delete(b.root, key)
}

func delete(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if root.key > key {
		root.left = delete(root.left, key)
	} else if root.key < key {
		root.right = delete(root.right, key)
	} else {
		if root.left == nil {
			return root.right
		}
		if root.right == nil {
			return root.left
		}
		x := min(root.right)
		x.right = deleteMin(root.right)
		x.left = root.left
		root = x
	}

	return root
}

func deleteMin(root *Node) *Node {
	if root.left == nil {
		return root.right
	}
	root.left = deleteMin(root.left)
	return root
}

func (b *Bst) get(key int) *Node {
	return get(b.root, key)
}

func get(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if root.key > key {
		return get(root.left, key)
	} else if root.key < key {
		return get(root.right, key)
	} else {
		return root
	}
}

func (b *Bst) inOrder() []*Node {
	result := &[]*Node{}
	inOrder(b.root, result)
	return *result
}

func inOrder(root *Node, result *[]*Node) {
	if root == nil {
		return
	}

	inOrder(root.left, result)
	*result = append(*result, root)
	inOrder(root.right, result)
}
