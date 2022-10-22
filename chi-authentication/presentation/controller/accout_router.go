package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateAccountController() *chi.Mux {
	ac := chi.NewRouter()

	ac.Post("/check_auth", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("認証されています"))
	})

	return ac
}
