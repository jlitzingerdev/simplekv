package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jlitzingerdev/simple-kv/kvdb"
)

type ReplyBody map[string]string

func configureServer() (*httptest.Server, *kvdb.Db) {
	db := kvdb.InitDb(&kvdb.DbConfig{})
	s := InitServer(db)
	ts := httptest.NewServer(s.router)
	return ts, db
}

func TestServerGet(t *testing.T) {
	ts, db := configureServer()
	db.Put([]byte("blech"), []byte("zab"))
	defer ts.Close()

	url := fmt.Sprintf("%s/v1/blech", ts.URL)
	res, err := http.Get(url)
	if err != nil {
		t.Errorf("Failed get: %v", err)
		t.FailNow()
	}

	reply, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("Read Body failed: %v", err)
		t.FailNow()
	}

	var data ReplyBody
	err = json.Unmarshal(reply, &data)
	if err != nil {
		t.Errorf("Unmarshal Body failed: %v", err)
		t.FailNow()
	}

	v, ok := data["blech"]
	if !ok {
		t.Errorf("data invalid %v", data)
		t.FailNow()
	}

	if string(v) != "zab" {
		t.Errorf("v != zab")
		t.FailNow()
	}
}

type testPostBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func TestPost(t *testing.T) {
	ts, db := configureServer()
	defer ts.Close()

	url := fmt.Sprintf("%s/v1/insert", ts.URL)

	body, err := json.Marshal(testPostBody{"bug", "buz"})
	if err != nil {
		t.Errorf("Failed marshal: %v", err)
		t.FailNow()
	}

	res, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		t.Errorf("Failed post: %v", err)
		t.FailNow()
	}

	if res.StatusCode != 200 {
		t.Errorf("StatusCode != 200")
		t.FailNow()
	}

	v := db.GetString("bug")
	if string(v) != "buz" {
		t.Errorf("v != buz")
		t.FailNow()
	}
}
