package week2

import "testing"
import "math/rand"
import "time"

func generateTree(nodeNum int) *Bst {
	tree := &Bst{}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < nodeNum; i++ {
		num := r1.Intn(20)
		tree.insert(num, num)
	}
	return tree
}
func TestBstInsert(t *testing.T) {
	tree := generateTree(20)
	inorder := tree.inOrder()
	t.Log(inorder)
}

func TestBstGet(t *testing.T) {
	tree := generateTree(100)
	inorder := tree.inOrder()
	t.Log(inorder)
	t.Log(time.Now())
	node := tree.get(3)
	t.Log(time.Now())
	t.Log(node)
}

func TestBstDelete(t *testing.T) {
	tree := generateTree(20)
	inorder := tree.inOrder()
	t.Log(inorder)
	tree.delete(3)
	inorder = tree.inOrder()
	t.Log(inorder)
}

func BenchmarkBst(b *testing.B) {
	var tree *Bst = generateTree(100000)
	for i := 0; i < b.N; i++ {
		tree.get(i)
	}
}
