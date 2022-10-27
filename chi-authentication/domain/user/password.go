package user

import (
	"chi_sample/common/utils"
	"errors"
)

type password struct {
	value string
}

// 新しいパスワードを作成するファクトリ関数
// セキュリティ的にuserモデルと分離してある
func NewPassword(pass string) (password, error) {
	if pass == "" {
		return password{}, errors.New("パスワードが空です。")
	}

	hashedPass := utils.CreateHash(pass)

	return password{value: hashedPass}, nil
}
