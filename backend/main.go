package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Msg string `json:"msg"`
}

func main() {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.HandleFunc("/", HandleHome)

	log.Fatal(server.ListenAndServe())
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	message := Message{Msg: "Welcome home!"}
	jaysawn, _ := json.Marshal(message)
	w.Write(jaysawn)
}
