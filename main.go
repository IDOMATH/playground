package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome home"))
}
