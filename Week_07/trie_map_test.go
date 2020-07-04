package week7

import "testing"

func TestTrieMapInsert(t *testing.T) {
	trie := NewNode()
	trie.Insert("app", 1)
	trie.Insert("apple", 2)
	t.Log(trie.next)

}

func TestTrieMapSearch(t *testing.T) {
	trie := NewNode()
	trie.Insert("app", 1)
	trie.Insert("apps", 2)
	x := trie.Search("app")
	t.Log(x)
}

func TestTrieMapKeyPrefix(t *testing.T) {
	trie := NewNode()
	trie.Insert("app", 1)
	trie.Insert("apple", 2)
	trie.Insert("apples", 3)
	trie.Insert("girls", 3)
	x := trie.KeyPrefix("")
	t.Log(x)
	x = trie.KeyPrefix("g")
	t.Log(x)
}
