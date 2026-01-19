package db

import "database/sql"

type BlogStore struct {
	DB *sql.DB
}

func NewBlogStore(db *sql.DB) *BlogStore {
	return &BlogStore{
		DB: db,
	}
}
