package user

type IUserRepository interface {
	Create(id, name, mail, imagePath, password string) error
}
