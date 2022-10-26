package user

import "errors"

type password struct {
	value string
}

func NewPassword(pass string) (password, error) {
	if pass == "" {
		return password{}, errors.New("パスワードが空です。")
	}
	return password{value: pass}, nil
}
