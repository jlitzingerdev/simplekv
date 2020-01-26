package kvdb_test

import (
	"github.com/jlitzingerdev/simple-kv/kvdb"
	"testing"
)

func TestPushPop(t *testing.T) {
	stack := kvdb.NewNodeStack()
	stack.Push(kvdb.NewStringNode("a", "foo"))
	stack.Push(kvdb.NewStringNode("b", "foo"))
	stack.Push(kvdb.NewStringNode("c", "foo"))

	n := stack.Pop()
	if string(n.Key()) != "c" {
		t.Errorf("n.Key() != c")
	}

	n = stack.Pop()
	if string(n.Key()) != "b" {
		t.Errorf("n.Key() != b")
	}

	n = stack.Pop()
	if string(n.Key()) != "a" {
		t.Errorf("n.Key() != a")
	}
}
