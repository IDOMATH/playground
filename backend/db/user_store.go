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

func (s *UserStore) LogIn(email, pass string) error {
	query := `SELECT email, password_hash FROM users WHERE email = ?`
}
