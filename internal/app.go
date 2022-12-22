package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	s, err := newServer("docdb.data", "8080")
	if err != nil {
		log.Fatal(err)
	}
	defer s.data.Close()

	s.reindex()

	router := httprouter.New()
	router.POST("/docs", s.addDocument)
	router.GET("/docs", s.searchDocuments)
	router.GET("/docs/:id", s.getDocument)

	log.Println("Listening on " + s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, router))
}
