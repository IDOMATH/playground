package db

import (
	"database/sql"

	"github.com/idomath/playground/backend/models"
	"golang.org/x/crypto/bcrypt"
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
	var user models.User
	query := `SELECT email, password_hash FROM users WHERE email = ?`
	err := s.DB.QueryRow(query, email).Scan(&user)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(pass))
	return err
}
