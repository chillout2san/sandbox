package service

import "chi_sample/domain/user"

// 既に会員登録されているかどうかをメールアドレスで判定する
func CheckRegistered(ur user.IUserRepository, mail string) bool {
	user, _ := ur.GetByMail(mail)

	return user.IsValid()
}
