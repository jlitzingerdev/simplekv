// Whitebox tests for tree.go
// This module intentionally tests only unexported functions in tree.go

package kvdb

import (
	"testing"
)

func checkNode(t *testing.T, n, parent *Node, expect *Node, msg string) {
	if n.Less(expect) {
		t.Errorf("%s: n.data != d1", msg)
	}

	if n.color != Red {
		t.Errorf("%s: n.color != Red", msg)
	}

	if n.parent != parent {
		t.Errorf("%s: n.parent != parent", msg)
	}
}

// doInsert adds a node respecting BST semantics whose color is
// always red.
func TestDoInsert(t *testing.T) {
	tree := NewTree()
	n := tree.doInsert(NewStringNode("l", "foo1"))

	if tree.root == nil {
		t.Errorf("Node not inserted")
		t.FailNow()
	}
	checkNode(t, tree.root, nil, n, "n")

	n2 := tree.doInsert(NewStringNode("q", "foo1"))

	checkNode(t, tree.root, nil, n, "n")
	checkNode(t, tree.root.right, n, n2, "n2")

	n3 := tree.doInsert(NewStringNode("d", "foo1"))

	checkNode(t, tree.root, nil, n, "n")
	checkNode(t, tree.root.right, n, n2, "n2")
	checkNode(t, tree.root.left, n, n3, "n3")

	n4 := tree.doInsert(NewStringNode("a", "foo1"))

	checkNode(t, tree.root, nil, n, "n")
	checkNode(t, tree.root.right, n, n2, "n2")
	l := tree.root.left
	checkNode(t, l, n, n3, "n3")
	checkNode(t, l.left, n3, n4, "n4")

	n5 := tree.doInsert(NewStringNode("r", "foo1"))

	checkNode(t, tree.root, nil, n, "n")
	r := tree.root.right
	checkNode(t, r, n, n2, "n2")
	checkNode(t, l, n, n3, "n3")
	checkNode(t, l.left, n3, n4, "n4")
	checkNode(t, r.right, n2, n5, "n5")
}

// If the inserted node is root, it's color is black
func TestColorRootBlack(t *testing.T) {
	tree := NewTree()
	n := tree.doInsert(NewStringNode("l", "foo1"))
	tree.reColor(n)
	if n.color != Black {
		t.Errorf("Root was not properly colored")
		t.FailNow()
	}
}

// If inserted node's parent is black, do nothing
func TestColorParentBlack(t *testing.T) {
	tree := NewTree()
	n := tree.doInsert(NewStringNode("l", "foo1"))
	tree.reColor(n)
	if n.color != Black {
		t.Errorf("Root was not properly colored")
		t.FailNow()
	}

	n = tree.doInsert(NewStringNode("a", "foo2"))
	tree.reColor(n)
	if n.color != Red {
		t.Errorf("New child was not properly colored")
		t.FailNow()
	}

	n = tree.doInsert(NewStringNode("m", "foo2"))
	tree.reColor(n)
	if n.color != Red {
		t.Errorf("New child was not properly colored")
		t.FailNow()
	}
}

func TestRedUncleLeftCaseNoRecurse(t *testing.T) {

	tree := NewTree()
	n1 := tree.doInsert(NewStringNode("l", "foo1"))
	tree.reColor(n1)

	n2 := tree.doInsert(NewStringNode("g", "foo2"))
	tree.reColor(n2)

	n3 := tree.doInsert(NewStringNode("p", "foo2"))
	tree.reColor(n3)

	n4 := tree.doInsert(NewStringNode("a", "foo2"))
	tree.reColor(n4)

	if n1.color != Black {
		t.Errorf("Root must be black")
		t.FailNow()
	}

	if n2.color != Black {
		t.Errorf("Parent should be black")
		t.FailNow()
	}

	if n3.color != Black {
		t.Errorf("Uncle should be black")
		t.FailNow()
	}

	if n4.color != Red {
		t.Errorf("Node should be red")
		t.FailNow()
	}
}

func TestNilUncleLeftCase(t *testing.T) {
	tree := NewTree()
	n1 := tree.doInsert(NewStringNode("l", "foo1"))
	tree.reColor(n1)

	n2 := tree.doInsert(NewStringNode("g", "foo2"))
	tree.reColor(n2)

	n4 := tree.doInsert(NewStringNode("a", "foo2"))
	tree.reColor(n4)

	if tree.root.color != Black {
		t.Errorf("Root must be black")
		t.FailNow()
	}

	if n2.color != Black {
		t.Errorf("Parent should be black")
		t.FailNow()
	}

	if n4.color != Red {
		t.Errorf("Node should be red")
		t.FailNow()
	}
}

func TestRedUncleLeftCaseRecurse(t *testing.T) {

	tree := NewTree()
	n1 := tree.doInsert(NewStringNode("l", "foo1"))
	tree.reColor(n1)

	n2 := tree.doInsert(NewStringNode("g", "foo2"))
	tree.reColor(n2)

	n3 := tree.doInsert(NewStringNode("s", "foo2"))
	tree.reColor(n3)

	n4 := tree.doInsert(NewStringNode("a", "foo2"))
	tree.reColor(n4)

	n5 := tree.doInsert(NewStringNode("o", "foo2"))
	tree.reColor(n5)

	n6 := tree.doInsert(NewStringNode("p", "foo2"))
	tree.reColor(n6)

	n7 := tree.doInsert(NewStringNode("n", "foo2"))
	tree.reColor(n7)

	n8 := tree.doInsert(NewStringNode("m", "foo2"))
	tree.reColor(n8)

	if n1.color != Black {
		t.Errorf("Root must be black")
		t.FailNow()
	}

	if n2.color != Black {
		t.Errorf("Parent should be black")
		t.FailNow()
	}

	if n3.color != Black {
		t.Errorf("Uncle should be black")
		t.FailNow()
	}

	if n4.color != Red {
		t.Errorf("Node should be red")
		t.FailNow()
	}
}

func TestGetRotateCase(t *testing.T) {
	n1 := NewStringNode("l", "bar")
	n2 := NewStringNode("h", "baz")
	n3 := NewStringNode("a", "foo")

	rcase := getRotateCase(n3, n2, n1)
	if rcase != LeftLeft {
		t.Errorf("rcase != LeftLeft")
		t.FailNow()
	}

	rcase = getRotateCase(n2, n3, n1)
	if rcase != LeftRight {
		t.Errorf("rcase != LeftRight")
		t.FailNow()
	}

	rcase = getRotateCase(n1, n2, n3)
	if rcase != RightRight {
		t.Errorf("rcase != RightRight")
		t.FailNow()
	}

	rcase = getRotateCase(n2, n1, n3)
	if rcase != RightLeft {
		t.Errorf("rcase != RightRight")
		t.FailNow()
	}
}

func TestRightRotateRoot(t *testing.T) {

	tree := NewTree()
	n1 := tree.doInsert(NewStringNode("l", "bar"))
	n2 := tree.doInsert(NewStringNode("h", "baz"))
	n3 := tree.doInsert(NewStringNode("i", "foo"))
	tree.doInsert(NewStringNode("a", "foo"))

	tree.rightRotate(n1)
	if n1.parent != n2 {
		t.Errorf("n1.parent != n2")
	}

	if tree.root != n2 {
		t.Errorf("tree.root != n2")
	}

	if n2.right != n1 {
		t.Errorf("n2.right != n1")
	}

	if n1.left != n3 {
		t.Errorf("n1.left != n3")
	}
}

func TestRightRotateNonRootLeftTree(t *testing.T) {
	tree := NewTree()
	n1 := tree.doInsert(NewStringNode("l", "bar"))
	n2 := tree.doInsert(NewStringNode("h", "baz"))
	n3 := tree.doInsert(NewStringNode("i", "foo"))
	n4 := tree.doInsert(NewStringNode("e", "foo"))
	n5 := tree.doInsert(NewStringNode("b", "foo"))
	n6 := tree.doInsert(NewStringNode("a", "foo"))
	n7 := tree.doInsert(NewStringNode("c", "foo"))

	tree.rightRotate(n2)
	if n1.left != n4 {
		t.Errorf("n1.left != n4")
	}

	if tree.root != n1 {
		t.Errorf("tree.root != n1")
	}

	if n4.right != n2 {
		t.Errorf("n2.right != n1")
	}

	if n2.right != n3 {
		t.Errorf("n1.left != n3")
	}

	if n4.left != n5 {
		t.Errorf("n4.left != n5")
	}

	if n5.left != n6 {
		t.Errorf("n4.left != n5")
	}

	if n5.right != n7 {
		t.Errorf("n4.left != n5")
	}
}

func TestRightRotateNonRootRightTree(t *testing.T) {

	tree := NewTree()
	n1 := tree.doInsert(NewStringNode("l", "bar"))
	n2 := tree.doInsert(NewStringNode("r", "baz"))
	n3 := tree.doInsert(NewStringNode("o", "foo"))
	n4 := tree.doInsert(NewStringNode("s", "foo"))

	tree.rightRotate(n2)
	if n1.right != n3 {
		t.Errorf("n1.left != n4")
	}

	if tree.root != n1 {
		t.Errorf("tree.root != n1")
	}

	if n3.right != n2 {
		t.Errorf("n2.right != n1")
	}

	if n2.right != n4 {
		t.Errorf("n1.left != n3")
	}

}

func TestLeftRotateNonRoot(t *testing.T) {

	tree := NewTree()
	n1 := tree.doInsert(NewStringNode("l", "bar"))
	n2 := tree.doInsert(NewStringNode("h", "baz"))
	n3 := tree.doInsert(NewStringNode("j", "foo"))
	n4 := tree.doInsert(NewStringNode("i", "foo"))

	if n1.left != n2 {
		t.Errorf("n1.left != n2")
		t.FailNow()
	}
	tree.leftRotate(n2)
	if tree.root != n1 {
		t.Errorf("tree.root != n1")
		t.FailNow()
	}

	if n1.left != n3 {
		t.Errorf("n1.left != n3")
		t.FailNow()
	}

	if n3.left != n2 {
		t.Errorf("n3.left != n2")
		t.FailNow()
	}

	if n2.right != n4 {
		t.Errorf("n2.right != n4")
		t.FailNow()
	}
}

func TestLeftRotateRoot(t *testing.T) {

	tree := NewTree()
	n1 := tree.doInsert(NewStringNode("l", "bar"))
	n2 := tree.doInsert(NewStringNode("o", "baz"))
	n3 := tree.doInsert(NewStringNode("j", "foo"))
	n4 := tree.doInsert(NewStringNode("m", "foo"))

	tree.leftRotate(n1)
	if tree.root != n2 {
		t.Errorf("tree.root != n2")
		t.FailNow()
	}

	if n1.right != n4 {
		t.Errorf("n1.right != n3")
		t.FailNow()
	}

	if n2.left != n1 {
		t.Errorf("n2.left != n1")
		t.FailNow()
	}

	if n1.left != n3 {
		t.Errorf("n1.left != n3")
		t.FailNow()
	}
}

func TestLeftRotateNonRootLeftChild(t *testing.T) {

	tree := NewTree()
	n1 := tree.doInsert(NewStringNode("l", "bar"))
	n2 := tree.doInsert(NewStringNode("d", "baz"))
	n3 := tree.doInsert(NewStringNode("c", "foo"))
	n4 := tree.doInsert(NewStringNode("f", "foo"))
	n5 := tree.doInsert(NewStringNode("a", "foo"))

	tree.leftRotate(n2)
	if n1.left != n4 {
		t.Errorf("n1.left != n4")
	}

	if n4.left != n2 {
		t.Errorf("n1.left != n4")
	}

	if n2.left != n3 {
		t.Errorf("n2.left != n3")
	}

	if n3.left != n5 {
		t.Errorf("n3.left != n5")
	}
}

func TestLeftRotateNonRootRightChild(t *testing.T) {

	tree := NewTree()
	n1 := tree.doInsert(NewStringNode("l", "bar"))
	n2 := tree.doInsert(NewStringNode("r", "baz"))
	n3 := tree.doInsert(NewStringNode("t", "foo"))
	n4 := tree.doInsert(NewStringNode("s", "foo"))

	tree.leftRotate(n2)
	if n1.right != n3 {
		t.Errorf("n1.left != n4")
	}

	if n3.left != n2 {
		t.Errorf("n3.left != n2")
	}

	if n2.right != n4 {
		t.Errorf("n2.right != n4")
	}
}

func TestLeftRightRoot(t *testing.T) {
	tree := NewTree()

	n1 := tree.doInsert(NewStringNode("l", "bar"))
	tree.reColor(n1)
	n2 := tree.doInsert(NewStringNode("r", "baz"))
	tree.reColor(n2)
	n2.color = Black
	n3 := tree.doInsert(NewStringNode("d", "foo"))
	tree.reColor(n3)
	n4 := tree.doInsert(NewStringNode("e", "foo"))
	tree.reColor(n4)

	if tree.root != n4 {
		t.Errorf("tree.root != n4")
	}

	if tree.root.color != Black {
		t.Errorf("tree.root.color != Black")
	}

	if n4.left != n3 {
		t.Errorf("n4.left != n3")
	}

	if n4.right != n1 {
		t.Errorf("n4.right != n1")
	}
}

func TestRightRightRoot(t *testing.T) {
	tree := NewTree()

	n1 := tree.doInsert(NewStringNode("l", "bar"))
	tree.reColor(n1)
	n2 := tree.doInsert(NewStringNode("d", "baz"))
	tree.reColor(n2)
	n2.color = Black
	n3 := tree.doInsert(NewStringNode("o", "foo"))
	tree.reColor(n3)
	n4 := tree.doInsert(NewStringNode("p", "foo"))
	tree.reColor(n4)

	if tree.root != n3 {
		t.Errorf("tree.root != n3")
	}

	if n3.right != n4 {
		t.Errorf("n3.right != n4")
	}

	if n3.left != n1 {
		t.Errorf("n3.left != n1")
	}
}

func TestRightLeftRoot(t *testing.T) {
	tree := NewTree()

	n1 := tree.doInsert(NewStringNode("l", "bar"))
	tree.reColor(n1)
	n2 := tree.doInsert(NewStringNode("d", "baz"))
	tree.reColor(n2)
	n2.color = Black
	n3 := tree.doInsert(NewStringNode("o", "foo"))
	tree.reColor(n3)
	n4 := tree.doInsert(NewStringNode("m", "foo"))
	tree.reColor(n4)

	if tree.root != n4 {
		t.Errorf("tree.root != n4")
	}

	if n4.right != n3 {
		t.Errorf("n4.right != n3")
	}
}
