package service

import (
	"chi_sample/domain/user"
	"context"
)

// 既に会員登録されているかどうかをメールアドレスで判定する
func CheckRegistered(ctx context.Context, ur user.IUserRepository, mail string) bool {
	user, _ := ur.GetByMail(ctx, mail)

	return user.IsValid()
}
