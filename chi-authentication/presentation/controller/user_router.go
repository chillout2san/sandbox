package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// userControllerを返却する
func CreateUserController() *chi.Mux {
	ac := chi.NewRouter()

	ac.Post("/check_auth", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("認証されています"))
	})

	return ac
}
