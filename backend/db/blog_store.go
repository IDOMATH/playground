package db

import (
	"database/sql"

	"github.com/idomath/playground/backend/models"
)

type BlogStore struct {
	DB *sql.DB
}

func NewBlogStore(db *sql.DB) *BlogStore {
	return &BlogStore{
		DB: db,
	}
}

func (s *BlogStore) InsertBlog(blog *models.Blog) (int, error) {
	return 0, nil
}
