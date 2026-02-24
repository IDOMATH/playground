package repository

import (
	"github.com/IDOMATH/session/memorystore"
	"github.com/idomath/playground/backend/db"
)

type Repository struct {
	BlogStore *db.BlogStore
	UserStore *db.UserStore
	Session   *memorystore.MemoryStore[string]
}

func NewRepository(bs *db.BlogStore, us *db.UserStore, ms *memorystore.MemoryStore[string]) Repository {
	return Repository{BlogStore: bs, UserStore: us, Session: ms}
}
