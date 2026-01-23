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

func (s *BlogStore) InsertBlog(blog *models.Blog) error {
	query := `INSERT INTO blogs (title, body, author_id) values (?, ?, ?)`
	_, err := s.DB.Exec(query, blog.Title, blog.Body, blog.AuthorId)
	return err
}

func (s *BlogStore) GetAllBlogs() ([]*models.Blog, error) {
	var blogs []*models.Blog

	return blogs, nil
}
