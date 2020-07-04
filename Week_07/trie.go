package week7

// TrieNode trie node
type TrieNode struct {
	Val  interface{}
	Next [26]*TrieNode
}

// Trie is 字典树
type Trie struct {
	root *TrieNode
}

// New xxx
func New() Trie {
	return Trie{}
}

func put(t *TrieNode, key string, val interface{}, d int) *TrieNode {
	if t == nil {
		t = &TrieNode{}
	}

	if d == len(key) {
		t.Val = val
		return t
	}
	i := key[d] - 'a'
	t.Next[i] = put(t.Next[i], key, val, d+1)
	return t
}

func (t *Trie) put(key string, val interface{}) {
	t.root = put(t.root, key, val, 0)
}

func get(node *TrieNode, key string, d int) *TrieNode {
	if node == nil {
		return nil
	}

	if d == len(key) {
		return node
	}

	return get(node.Next[key[d]-'a'], key, d+1)
}

func (t *Trie) get(key string) interface{} {
	node := get(t.root, key, 0)
	if node == nil {
		return nil
	}
	return node.Val
}

func collect(node *TrieNode, pre string, q *[]string) {
	if node == nil {
		return
	}

	if node.Val != nil {
		*q = append(*q, pre)
	}

	for i := 0; i < len(node.Next); i++ {
		collect(node.Next[i], pre+string([]byte{byte(i) + 'a'}), q)
	}
}

func (t *Trie) keyPrefix(pre string) []string {
	q := []string{}
	node := get(t.root, pre, 0)
	collect(node, pre, &q)
	return q
}

func delete(node *TrieNode, key string, d int) *TrieNode {
	if node == nil {
		return nil
	}

	if d == len(key) {
		node.Val = nil
	} else {
		i := key[d] - 'a'
		node.Next[i] = delete(node.Next[i], key, d+1)
	}

	if node.Val != nil {
		return node
	}

	for i := 0; i < len(node.Next); i++ {
		if node.Next[i] != nil {
			return node
		}
	}
	return nil

}

func (t *Trie) delete(key string) {
	t.root = delete(t.root, key, 0)
}

func hasWord(root *TrieNode) bool {
	if root == nil {
		return false
	}

	if root.Val != nil {
		return true
	}
	for _, node := range root.Next {
		if hasWord(node) {
			return true
		}
	}
	return false
}

// StartsWith prefix
func (t *Trie) StartsWith(prefix string) bool {
	root := get(t.root, prefix, 0)
	if root == nil {
		return false
	}
	return hasWord(root)
}
