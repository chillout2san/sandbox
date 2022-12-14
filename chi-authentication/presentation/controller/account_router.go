package controller

import (
	"chi_sample/common/utils"
	"chi_sample/infrastructure"
	"chi_sample/infrastructure/repository/user"
	"chi_sample/presentation/middleware"
	"chi_sample/usecase/account/login"
	"chi_sample/usecase/account/register"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var ur = user.NewUserRepository(infrastructure.Db)
var ri = register.NewRegisterInteractor(ur)
var li = login.NewLoginInteractor(ur)

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

		result := ri.Interact(r.Context(), inputDto)
		res, _ := json.Marshal(result)
		w.Write(res)
	})

	ac.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		var inputDto login.InputDto

		err := middleware.MapInputDto(r, &inputDto)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res, _ := json.Marshal(login.OutputDto{
				Id:         "",
				Token:      "",
				ErrMessage: err.Error(),
			})
			w.Write(res)
			return
		}

		result := li.Interact(r.Context(), inputDto)
		res, _ := json.Marshal(result)
		w.Write(res)
	})

	ac.Post("/check_auth", func(w http.ResponseWriter, r *http.Request) {
		type OutputDto struct {
			HasAuth    bool   `json:"hasAuth"`
			ErrMessage string `json:"errMessage"`
		}

		var inputDto struct {
			Id    string `json:"id"`
			Token string `json:"token"`
		}

		err := middleware.MapInputDto(r, &inputDto)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res, _ := json.Marshal(OutputDto{
				HasAuth:    false,
				ErrMessage: err.Error(),
			})
			w.Write(res)
			return
		}

		err = utils.CheckJwt(inputDto.Id, inputDto.Token)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res, _ := json.Marshal(OutputDto{
				HasAuth:    false,
				ErrMessage: err.Error(),
			})
			w.Write(res)
			return
		}

		res, _ := json.Marshal(OutputDto{
			HasAuth:    true,
			ErrMessage: "",
		})

		w.Write(res)
	})

	return ac
}
