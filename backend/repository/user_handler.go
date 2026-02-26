package repository

import (
	"encoding/json"
	"net/http"

	"github.com/idomath/playground/backend/models"
)

func (repo *Repository) HandleNewUser(w http.ResponseWriter, r *http.Request) {

	var user *models.User
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// log.Error("HandlePostNewUser", "error decoding user to json from body", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("getting user from body"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Inserted new user"))
}

func (repo *Repository) HandleLogin(w http.ResponseWriter, r *http.Request) {

}

func (repo *Repository) HandleLogout(w http.ResponseWriter, r *http.Request) {}
