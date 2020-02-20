package kvdb_test

import (
	"fmt"
	"github.com/jlitzingerdev/simple-kv/kvdb"
	"testing"
)

func TestInOrder(t *testing.T) {
	tr := kvdb.NewTree()
	tr.Insert([]byte("z"), []byte("foo"))
	tr.Insert([]byte("d"), []byte("foo"))
	tr.Insert([]byte("g"), []byte("foo"))
	tr.Insert([]byte("c"), []byte("foo"))
	tr.Insert([]byte("b"), []byte("foo"))
	tr.Insert([]byte("y"), []byte("foo"))
	tr.Insert([]byte("a"), []byte("foo"))

	results := []string{}
	tr.InOrder(func(node *kvdb.Node) {
		results = append(results, string(node.Key()))
	})

	last := results[0]
	for _, v := range results[1:] {
		if last > v {
			t.Errorf("%s not less than %s", last, v)
			t.FailNow()
		}
		last = v
	}
}

func TestNoDuplicates(t *testing.T) {
	tr := kvdb.NewTree()
	tr.Insert([]byte("a"), []byte("bar"))
	tr.Insert([]byte("a"), []byte("foo"))
	results := []string{}
	tr.InOrder(func(node *kvdb.Node) {
		results = append(results, string(node.Key()))
	})
	if len(results) != 1 {
		t.Errorf("tree allows duplicates")
		t.FailNow()
	}
}

func TestGet(t *testing.T) {
	tr := kvdb.NewTree()
	tr.Insert([]byte("a"), []byte("foo"))
	v := tr.Get([]byte("a"))
	if string(v) != "foo" {
		t.Errorf("v != foo")
		t.FailNow()
	}

	v = tr.Get([]byte("b"))
	if v != nil {
		t.Errorf("v != nil")
		t.FailNow()
	}
}

func TestDelete(t *testing.T) {
	tr := kvdb.NewTree()
	tr.Insert([]byte("h"), []byte("wtf"))

	v := tr.Get([]byte("h"))
	if string(v) != "wtf" {
		t.Errorf("v != wtf")
		t.FailNow()
	}

	tr.Delete([]byte("h"))

	v = tr.Get([]byte("h"))
	if v != nil {
		t.Errorf("v != nil")
		t.FailNow()
	}
}

func TestDeleteNil(t *testing.T) {
	tr := kvdb.NewTree()
	tr.Delete([]byte("h"))
}

func TestPutGet(t *testing.T) {
	tr := kvdb.NewTree()
	tr.Insert([]byte("foo"), []byte("bar"))
	tr.Insert([]byte("biz"), []byte("baz"))

	tr.InOrder(func(node *kvdb.Node) {
		fmt.Printf("%s, %s\n", string(node.Key()), string(node.Value()))
	})

	v := tr.Get([]byte("foo"))
	if string(v) != "bar" {
		t.Errorf("%v != bar", v)
		t.FailNow()
	}

	v = tr.Get([]byte("biz"))
	if string(v) != "baz" {
		t.Errorf("%v != baz", v)
		t.FailNow()
	}
}
