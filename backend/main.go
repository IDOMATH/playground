package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/idomath/playground/backend/db"
	"github.com/idomath/playground/backend/repository"
)

type Message struct {
	Msg string `json:"msg"`
}

func main() {
	log.Fatal(run())
}

func run() error {

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	dbHost := "localhost"
	dbPort := "5432"
	dbName := "portfolio"
	dbUser := "postgres"
	dbPass := "postgres"
	dbSsl := "disable"

	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", dbHost, dbPort, dbName, dbUser, dbPass, dbSsl)
	fmt.Println("Connecting to Postgres")
	postgresDb, err := db.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.Repository{BlogStore: *db.NewBlogStore(postgresDb.SQL)}

	addRoutes(router, &repo)

	return server.ListenAndServe()
}

func addRoutes(router *http.ServeMux, repo *repository.Repository) {

	router.HandleFunc("/", HandleHome)
	router.HandleFunc("POST /blogs/", repo.HandleCreateBlogPost)
	router.HandleFunc("GET /blogs", repo.HandleGetBlogPosts)
	router.HandleFunc("GET /blogs/{id}", repo.HandleGetBlogById)
	router.HandleFunc("DELETE /blogs/{id}", repo.HandleDeleteBlogPost)

}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	message := Message{Msg: "Welcome home!"}
	jaysawn, _ := json.Marshal(message)
	w.Write(jaysawn)
}
