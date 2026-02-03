package db

import (
	"fmt"
	"testing"

	"github.com/idomath/playground/backend/db"
)

func TestBlogStore_ConnectToDatabase(t *testing.T) {
	dbHost := "localhost"
	dbPort := "5432"
	dbName := "portfolio"
	dbUser := "postgres"
	dbPass := "postgres"
	dbSsl := "disable"

	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", dbHost, dbPort, dbName, dbUser, dbPass, dbSsl)
	fmt.Println("Connecting to Postgres")
	_, err := db.ConnectSQL(connectionString)
	if err != nil {
		t.Error("error connecting to postgres")
	}
}
