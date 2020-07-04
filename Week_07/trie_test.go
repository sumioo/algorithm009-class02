package week7

import "testing"

func TestTriePut(t *testing.T) {
	trie := New()
	trie.put("apple", 1)
	trie.put("app", "a")
	trie.put("application", 1)

	pres := trie.keyPrefix("")
	// t.Log(trie.root)
	for i := range pres {
		t.Log(pres[i])
	}

	// t.Logf("%T %v", v2, v2.(int) == int)
}

func TestTrieGet(t *testing.T) {
	trie := New()
	trie.put("apple", 1)
	trie.put("app", "a")
	v := trie.get("app")
	t.Logf("%T %v", v, v)

	v2 := trie.get("apple")
	if _, ok := v2.(int); ok {
		t.Log("S")
	}
}

func TestTrieDelete(t *testing.T) {
	trie := New()
	trie.put("apple", 1)
	trie.put("app", "a")
	trie.delete("apple")
	t.Log(trie.keyPrefix(""))
	trie.delete("app")
	t.Log(trie.keyPrefix(""))
	t.Log(trie.root)
}
