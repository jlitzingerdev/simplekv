package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v4"
	"github.com/jlitzingerdev/simple-kv/kvdb"
)

type Server struct {
	db     *kvdb.Db
	router *chi.Mux
}

// Handler for GET /v1/{key}.  Returns a JSON object of the form
// {"key": "value"}.
func (s *Server) GetKey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		k := chi.URLParam(r, "key")
		v := s.db.GetString(strings.TrimSpace(k))
		if v == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// This is *not* going to work as desired for non
		// string data
		body := map[string]string{}
		body[k] = string(v)
		blob, err := json.Marshal(body)
		if err != nil {
			fmt.Println("Failed encoding ", err)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(blob)
	}
}

type PostBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s *Server) PostKey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)

		var body PostBody

		err := dec.Decode(&body)
		if err != nil {
			fmt.Println("Bad data ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Printf("%s = %s", body.Key, body.Value)
		s.db.Put([]byte(body.Key), []byte(body.Value))
		w.WriteHeader(http.StatusOK)
	}

}

// Start the server, runs until exit is called or program
// exits
func (s *Server) StartServer() {
	http.ListenAndServe(":10000", s.router)
}

func InitServer(db *kvdb.Db) *Server {
	s := &Server{db, chi.NewRouter()}
	s.router.Route("/v1", func(r chi.Router) {
		r.Get("/{key}", s.GetKey())
		r.Post("/insert", s.PostKey())
	})
	return s
}
