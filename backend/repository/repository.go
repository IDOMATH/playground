package repository

import "github.com/idomath/playground/backend/db"

type Repository struct {
	BlogStore db.BlogStore
}

func NewRepository(bs db.Blogstore) Repository {
	return Repository{Blogstore: bs}
}
