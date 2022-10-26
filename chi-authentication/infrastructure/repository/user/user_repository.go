package user

type userRepository struct {
}

func NewAccountRepository() userRepository {
	return userRepository{}
}

func (a userRepository) Create(id, name, mail, imagePath, password string) {

}

func (a userRepository) Read() {

}
