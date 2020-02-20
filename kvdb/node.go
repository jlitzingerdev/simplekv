package kvdb

import (
	"bytes"
	"time"
)

// Simple node for an arbitrary key/value pair.  Timestamps are epoch
// time.
type Node struct {
	key       []byte
	value     []byte
	color     Color
	left      *Node
	right     *Node
	parent    *Node
	timestamp int64
	tombstone bool
}

func NewNode(key, value []byte) *Node {
	n := &Node{key, value, Red, nil, nil, nil, -1, false}
	n.timestamp = time.Now().Unix()
	return n
}

func NewStringNode(key, value string) *Node {
	n := &Node{[]byte(key), []byte(value), Red, nil, nil, nil, -1, false}
	n.timestamp = time.Now().Unix()
	return n
}

func (n *Node) Key() []byte {
	return n.key
}

func (n *Node) Value() []byte {
	return n.value
}

func (n *Node) Timestamp() int64 {
	return n.timestamp
}

// Returns -1, 0, or 1 depending on whether lhs.key is less than, equal to,
// or greater than rhs.key
func (lhs *Node) Compare(rhs *Node) int {
	return bytes.Compare(lhs.key, rhs.key)
}

func (lhs *Node) Less(rhs *Node) bool {
	r := lhs.Compare(rhs)
	return r == -1
}

// Returns -1, 0, or 1 depending on whether lhs is less than, equal to,
// or greater than key
func (lhs *Node) CompareKey(key []byte) int {
	return bytes.Compare(lhs.key, key)
}

func (n *Node) SetValue(value []byte) {
	n.value = value
	n.timestamp = time.Now().Unix()
	n.tombstone = false
}

func (n *Node) Delete() {
	n.tombstone = true
}
