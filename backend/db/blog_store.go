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

	query := `SELECT title, body, author_id FROM blogs`

	rows, err := s.DB.Query(query)
	if err != nil {
		return blogs, err
	}

	defer rows.Close()
	for rows.Next() {
		var row models.Blog
		err := rows.Scan(
			row.Title,
			row.Body,
			row.AuthorId)

		if err != nil {
			return blogs, err
		}
		blogs = append(blogs, &row)
	}

	return blogs, nil
}

func (s *BlogStore) GetBlogById(id int) (models.Blog, error) {
	var blog models.Blog

	query := `SELECT title, body, author_id FROM blogs WHERE id = ?`

	err := s.DB.QueryRow(query, id).Scan(blog.Title, blog.Body, blog.AuthorId)

	return blog, err
}

func (s *BlogStore) DeleteBlog(id int) error {
	query := `DELETE FROM blogs WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
