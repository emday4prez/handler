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
	log.Print("saving")
	return nil
}

func main() {
	theDb := &SimpleDB{}
	server := &Server{db: theDb}

	http.HandleFunc("/hello", server.HelloHandler)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s *Server) HelloHandler(w http.ResponseWriter, r *http.Request) {
	err := s.db.Save("hello, world")
	if err != nil {
		http.Error(w, "Could not save data", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Data saved: hello, world")
}
