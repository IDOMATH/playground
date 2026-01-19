package repository

import "net/http"

func (repo *Repository) HandleCreateBlogPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
