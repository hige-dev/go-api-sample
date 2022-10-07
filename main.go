package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    ":8888",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Printf("error: %+v", err)
	}
}
