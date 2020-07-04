package week7

// Node trie tree node
type Node struct {
	next map[rune]*Node
	val  interface{}
}

//NewNode create a Node
func NewNode() *Node {
	n := &Node{}
	n.next = make(map[rune]*Node)
	return n
}

// Insert a trie key
func (n *Node) Insert(key string, val interface{}) {
	curr := n
	for _, r := range key {
		next, ok := curr.next[r]
		if !ok {
			next = NewNode()
			curr.next[r] = next
		}
		curr = next
	}
	curr.val = val
	return
}

// Search trie tree
func (n *Node) Search(key string) interface{} {
	node := n.get(key)
	if node == nil {
		return nil
	}
	return node.val
}

func (n *Node) get(key string) *Node {
	curr := n
	for _, r := range key {
		next, ok := curr.next[r]
		if !ok {
			return nil
		}
		curr = next
	}
	return curr
}

func collectR(root *Node, prefix string, q *[]string) {
	if root == nil {
		return
	}

	if root.val != nil {
		*q = append(*q, prefix)
	}

	for r, node := range root.next {
		collectR(node, prefix+string([]rune{r}), q)
	}
}

// KeyPrefix find key whit prefix
func (n *Node) KeyPrefix(prefix string) []string {
	curr := n.get(prefix)
	if curr == nil {
		return nil
	}
	q := []string{}
	collectR(curr, prefix, &q)
	return q
}


