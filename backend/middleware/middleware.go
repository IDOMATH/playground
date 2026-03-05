package middleware

import (
	"fmt"
	"net/http"

	"github.com/idomath/playground/backend/repository"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Authorize(repo *repository.Repository) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			_, found := repo.Session.Get(r.Header.Get("cheetauth"))
			if !found {
				fmt.Println("NOT AUTHENTICATED")
				// Potentially do some rerouting if the endpoint is protected
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(r.Header.Get("cheetauth")))
				return
			}

			w.Write([]byte("Authorized"))

			next(w, r)
		}
	}
}
