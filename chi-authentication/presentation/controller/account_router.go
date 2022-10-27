package controller

import (
	"chi_sample/infrastructure/repository/user"
	"chi_sample/presentation/middleware"
	"chi_sample/usecase/account/register"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var ur = user.NewUserRepository()
var ri = register.NewAccountInteractor(ur)

// accountControllerを返却する
func NewAccountController() *chi.Mux {
	ac := chi.NewRouter()

	ac.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		var inputDto register.InputDto

		err := middleware.MapInputDto(r, &inputDto)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res, _ := json.Marshal(register.OutputDto{
				IsRegistered: false,
				ErrMessage:   err.Error(),
			})
			w.Write(res)
			return
		}

		result := ri.Interact(inputDto)
		res, _ := json.Marshal(result)
		w.Write(res)
	})

	return ac
}
