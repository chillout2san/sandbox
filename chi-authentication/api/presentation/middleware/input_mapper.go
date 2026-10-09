package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func MapInputDto[T interface{}](r *http.Request, inputData *T) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("Can't read request body:%v;", err)
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, inputData)
	if err != nil {
		return fmt.Errorf("Can't parse request body:%v;", err)
	}

	return nil
}
