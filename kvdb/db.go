package kvdb

// Main database interface for simple-kv
import (
	"sync"
)

type DbConfig struct {
}

type Db struct {
	tree *Tree
	lock sync.Mutex
}

func (db *Db) GetString(key string) []byte {
	db.lock.Lock()
	defer db.lock.Unlock()
	return db.tree.Get([]byte(key))
}

func (db *Db) Put(key, value []byte) {
	db.lock.Lock()
	defer db.lock.Unlock()
	db.tree.Insert(key, value)
}

func (db *Db) Delete(key []byte) {
	db.lock.Lock()
	defer db.lock.Unlock()
	db.tree.Delete(key)
}

func InitDb(_ *DbConfig) *Db {
	return &Db{tree: NewTree()}
}
