package controller

import (
	"chi_sample/common/utils"
	"chi_sample/config"
	"chi_sample/infrastructure"
	"chi_sample/infrastructure/repository/user"
	"chi_sample/presentation/middleware"
	"chi_sample/usecase/account/login"
	"chi_sample/usecase/account/register"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

var ur = user.NewUserRepository(infrastructure.Db)
var ru = register.NewRegisterUseCase(ur)
var lu = login.NewLoginUseCase(ur)

// accountControllerを返却する
func NewAccountController() *chi.Mux {
	ac := chi.NewRouter()

	ac.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		var inputDto register.InputDto

		err := middleware.MapInputDto(r, &inputDto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res, _ := json.Marshal(map[string]interface{}{
				"isRegistered": false,
				"errMessage":   "リクエストに誤りがあります。",
			})
			w.Write(res)
			return
		}

		result := ru.Execute(r.Context(), inputDto)

		if result.ErrMessage != "" {
			w.WriteHeader(http.StatusInternalServerError)
			res, _ := json.Marshal(map[string]interface{}{
				"isRegistered": false,
				"errMessage":   result.ErrMessage,
			})
			w.Write(res)
			return
		}

		res, _ := json.Marshal(map[string]interface{}{
			"isRegistered": true,
		})

		w.Write(res)
	})

	ac.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		var inputDto login.InputDto

		err := middleware.MapInputDto(r, &inputDto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res, _ := json.Marshal(map[string]interface{}{
				"id":         "",
				"errMessage": "リクエストに誤りがあります。",
			})
			w.Write(res)
			return
		}

		result := lu.Execute(r.Context(), inputDto)
		if result.ErrMessage != "" {
			w.WriteHeader(http.StatusInternalServerError)
			res, _ := json.Marshal(map[string]interface{}{
				"id":         "",
				"errMessage": err.Error(),
			})
			w.Write(res)
			return
		}

		cookie := &http.Cookie{
			Name:     "token",
			Value:    result.Token,
			Path:     "/",
			Expires:  time.Now().Add(10 * time.Minute),
			Secure:   config.Environment.COOKIE_SECURE,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, cookie)

		res, _ := json.Marshal(map[string]interface{}{
			"id": result.Id,
		})

		w.Write(res)
	})

	ac.Post("/check_auth", func(w http.ResponseWriter, r *http.Request) {
		var inputDto struct {
			Id string `json:"id"`
		}

		err := middleware.MapInputDto(r, &inputDto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res, _ := json.Marshal(map[string]interface{}{
				"hasAuth":    false,
				"errMessage": "リクエストに誤りがあります。",
			})
			w.Write(res)
			return
		}

		token, err := r.Cookie("token")
		if err != nil {
			log.Println("check_auth failed:", err)
			w.WriteHeader(http.StatusBadRequest)
			res, _ := json.Marshal(map[string]interface{}{
				"hasAuth":    false,
				"errMessage": "トークンが確認できません。",
			})
			w.Write(res)
			return
		}

		err = utils.CheckJwt(inputDto.Id, token.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			res, _ := json.Marshal(map[string]interface{}{
				"hasAuth":    false,
				"errMessage": err.Error(),
			})
			w.Write(res)
			return
		}

		res, _ := json.Marshal(map[string]interface{}{
			"hasAuth": true,
		})

		w.Write(res)
	})

	return ac
}
