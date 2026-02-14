package repository

import "github.com/idomath/playground/backend/db"

type Repository struct {
	BlogStore *db.BlogStore
}

func NewRepository(bs *db.BlogStore) Repository {
	return Repository{BlogStore: bs}
}
