package todo

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct{}

func (s *Server) Start(port string) {
	// using default http handler: DefaultServeMux
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/todo/", todoIndexHandler)

	fmt.Println("listening on: ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi! This is a simple todo application"))
}
