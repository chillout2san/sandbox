package service

import (
	"chi_sample/domain/user"
	"context"
)

// 既に会員登録されているかどうかをメールアドレスで判定する
func CheckRegistered(ctx context.Context, ur user.IUserRepository, mail string) (bool, error) {
	user, err := ur.GetByMail(ctx, mail)

	if err != nil {
		return false, err
	}

	if user.Id() == "" {
		return false, nil
	}

	return true, nil
}
