package user

import (
	"chi_sample/common/utils"
	"chi_sample/infrastructure"
	"context"
	"fmt"
)

type userRepository struct {
}

func NewAccountRepository() userRepository {
	return userRepository{}
}

// 新しくユーザーを登録する。
func (a userRepository) Create(id, name, mail, imagePath, password string) error {
	sql := `INSERT INTO users(id, name, mail, imagePath, password) VALUE(?,?,?,?,?)`

	hashedPassword := utils.CreateHash(password)

	_, err := infrastructure.Db.ExecContext(context.TODO(), sql, id, mail, imagePath, hashedPassword)

	if err != nil {
		return fmt.Errorf("infra/create:%v", err)
	}

	return nil
}

func (a userRepository) Read() {

}
