// Copyright 2020 Jason Litzinger

// Use of this software is governed be the MIT license in
// LICENSE.txt.

// This module houses the implementation of the red-black tree housed in
// in memory as the inital store of keys as well as the SS table index
// implementations.
// Properties
// 1. Every node has a color -- red or black
// 2. The root of the tree is always black
// 3. There are no two adjacent red nodes
// 4. Every path from a node, to any of its decendent NULL leaves
//    has the same number of black nodes

package kvdb

import (
	"sync"
)

type Color int

const (
	Red Color = iota
	Black
)

type RotateCase int

const (
	LeftLeft RotateCase = iota
	LeftRight
	RightRight
	RightLeft
)

type Tree struct {
	root *Node
	lock sync.Mutex
}

func NewTree() *Tree {
	return &Tree{}
}

func (tree *Tree) getNode(key []byte) *Node {
	target := tree.root
	for target != nil {
		c := target.CompareKey(key)
		if c == -1 {
			target = target.left
		} else if c == 0 && !target.tombstone {
			return target
		} else {
			target = target.right
		}
	}
	return nil
}

func (tree *Tree) doInsert(n *Node) *Node {
	target := &tree.root
	parent := tree.root
	for *target != nil {
		parent = *target
		c := n.Compare(*target)
		if c == -1 {
			target = &((*target).left)
		} else if c == 0 {
			(*target).SetValue(n.value)
			return *target
		} else {
			target = &((*target).right)
		}
	}
	n.parent = parent
	*target = n
	return *target
}

// Right rotate tree rooted at oldRoot such that:
// newNode.right is set to oldRoot
func (tree *Tree) rightRotate(oldRoot *Node) {

	if oldRoot.left == nil {
		panic("Attempt to right rotate with nil left node")
	}

	leftChild := oldRoot.left
	leftChild.parent = oldRoot.parent
	tmp := leftChild.right
	leftChild.right = oldRoot

	oldRoot.left = tmp

	if oldRoot.parent == nil {
		tree.root = leftChild
	} else if oldRoot.parent.left == oldRoot {
		oldRoot.parent.left = leftChild
	} else {
		oldRoot.parent.right = leftChild
	}
	oldRoot.parent = leftChild
}

func (tree *Tree) leftRotate(oldRoot *Node) {
	if oldRoot.right == nil {
		panic("Attempt to left rotate with nil right node")
	}

	rightChild := oldRoot.right
	rightChild.parent = oldRoot.parent
	tmp := rightChild.left
	rightChild.left = oldRoot
	oldRoot.right = tmp
	if oldRoot.parent == nil {
		tree.root = rightChild
	} else if oldRoot.parent.left == oldRoot {
		oldRoot.parent.left = rightChild
	} else {
		oldRoot.parent.right = rightChild
	}
	oldRoot.parent = rightChild
}

func getRotateCase(newNode, parent, grandparent *Node) RotateCase {
	if parent.Less(grandparent) {
		if newNode.Less(parent) {
			return LeftLeft
		} else {
			return LeftRight
		}
	} else {
		if parent.Less(newNode) {
			return RightRight
		} else {
			return RightLeft
		}
	}
}

func (tree *Tree) reColor(newNode *Node) {
	target := newNode

	for target != nil {
		if target.parent == nil {
			if target.color != Black {
				newNode.color = Black
			}
			return
		}

		if target.parent.color == Black {
			return
		}

		parent := target.parent
		grandparent := parent.parent
		var uncle *Node
		if parent == grandparent.left {
			uncle = grandparent.right
		} else {
			uncle = grandparent.left
		}

		if uncle != nil && uncle.color == Red {
			parent.color = Black
			uncle.color = Black
			target = grandparent
		} else {
			rotateCase := getRotateCase(newNode, parent, grandparent)
			switch rotateCase {
			case LeftLeft:
				tree.rightRotate(grandparent)
				parent.color, grandparent.color = grandparent.color, parent.color
			case LeftRight:
				tree.leftRotate(parent)
				tree.rightRotate(grandparent)
				target.color, grandparent.color = grandparent.color, target.color
			case RightRight:
				tree.leftRotate(grandparent)
				parent.color, grandparent.color = grandparent.color, parent.color
			case RightLeft:
				tree.rightRotate(parent)
				tree.leftRotate(grandparent)
				target.color, grandparent.color = grandparent.color, target.color
			default:
				panic("No valid rotation, should not happen")
			}
			target = nil
		}
	}

}

func (tree *Tree) Insert(key, value []byte) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	n := NewNode(key, value)
	n = tree.doInsert(n)
	tree.reColor(n)
}

func (tree *Tree) Delete(key []byte) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	n := tree.getNode(key)
	if n != nil {
		n.Delete()
	}
}

func (tree *Tree) Get(key []byte) []byte {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	n := tree.getNode(key)
	if n != nil {
		return n.value
	}
	return nil
}

type TraversalOperation func(node *Node)

func (tree *Tree) InOrder(op TraversalOperation) {
	n := tree.root
	stack := NewNodeStack()

	for n != nil || !stack.Empty() {
		for n != nil {
			stack.Push(n)
			n = n.left
		}
		n = stack.Pop()
		op(n)
		n = n.right
	}
}
