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

func HandleHome(w http.ResponseWriter, r *http.Request) { // Set the origin to * to allow all origins, or specify a particular origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	message := Message{Msg: "Welcome home!"}
	jaysawn, _ := json.Marshal(message)
	w.Write(jaysawn)
}
