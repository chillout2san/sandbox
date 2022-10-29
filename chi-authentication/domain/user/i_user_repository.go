package user

type IUserRepository interface {
	Create(u User, p Password) error
	GetByMail(value string) (User, error)
	GetPassByMail(value string) (Password, error)
}
