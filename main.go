package main

import (
	"github.com/jlitzingerdev/simple-kv/api"
	"github.com/jlitzingerdev/simple-kv/kvdb"
)

func main() {
	db := kvdb.InitDb(&kvdb.DbConfig{})
	s := api.InitServer(db)
	s.StartServer()
}
