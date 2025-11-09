package main

import (
	"fmt"
	"log"
	"net/http"
)

type Datastore interface{ Save(data string) error }
type Server struct{ db Datastore }
type SimpleDB struct{}

func (db *SimpleDB) Save(data string) error {
	log.Printf("REAL_DB: Saving data: %s", data)

	return nil
}
func main() {
	realDB := &SimpleDB{}

	server := &Server{db: realDB}

	http.HandleFunc("/hello", server.HelloHandler)
}

func (s *Server) HelloHandler(w http.ResponseWriter, r *http.Request) {
	err := s.db.Save("hello, world")
	if err != nil {
		http.Error(w, "Could not save data", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Data saved: hello, world")
}
