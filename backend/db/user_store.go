package db

import (
	"database/sql"

	"github.com/idomath/playground/backend/models"
)

type UserStore struct {
	DB *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		DB: db,
	}
}

func (s *UserStore) InsertUser(user *models.User) error {
	query := `INSERT INTO users (email, password_hash) values (?, ?)`
	_, err := s.DB.Exec(query, user.Email, user.PasswordHash)
	return err
}
