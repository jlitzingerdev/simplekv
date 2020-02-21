package kvdb

// Main database interface for simple-kv

type DbConfig struct {
}

type Db struct {
	tree *Tree
}

func (db *Db) GetString(key string) []byte {
	return db.tree.Get([]byte(key))
}

func (db *Db) Put(key, value []byte) {
	db.tree.Insert(key, value)
}

func (db *Db) Delete(key []byte) {
	db.tree.Delete(key)
}

func InitDb(_ *DbConfig) *Db {
	return &Db{NewTree()}
}
