package repository

import (
	"net/http"
	"strconv"
)

func (repo *Repository) HandleCreateBlogPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte("error parsing form"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.PostForm.Get("title")
	r.PostForm.Get("body")
	r.PostForm.Get("author")

	w.WriteHeader(http.StatusCreated)
}

func (repo *Repository) HandleGetBlogPosts(w http.ResponseWriter, r *http.Request) {
	blogs, err := repo.BlogStore.GetAllBlogs()
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, blog := range blogs {
		w.Write([]byte(blog.Title))
	}
	w.WriteHeader(http.StatusOK)
}

func (repo *Repository) HandleGetBlogById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	blog, err := repo.BlogStore.GetBlogById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(blog.Title))

	w.WriteHeader(http.StatusOK)
}
