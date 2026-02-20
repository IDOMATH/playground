package repository

import "github.com/idomath/playground/backend/db"

type Repository struct {
	BlogStore *db.BlogStore
	UserStore *db.UserStore
}

func NewRepository(bs *db.BlogStore, us *db.UserStore) Repository {
	return Repository{BlogStore: bs, UserStore: us}
}
