package middleware

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func MapInputDto[T interface{}](r *http.Request, inputData *T) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Can't read request body:", err)
		return errors.New("入力した値に誤りがあります。")
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, inputData)
	if err != nil {
		log.Println("Can't parse request body:", err)
		return errors.New("入力した値に誤りがあります。")
	}
	return nil
}
