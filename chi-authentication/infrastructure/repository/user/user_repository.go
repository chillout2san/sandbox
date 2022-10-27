package user

import (
	"chi_sample/infrastructure"
	"context"
	"fmt"
)

type userRepository struct {
}

// ユーザーリポジトリを返却する
func NewUserRepository() userRepository {
	return userRepository{}
}

// 新しくユーザーを登録する。
func (a userRepository) Create(id, name, mail, imagePath, password string) error {
	sql := `INSERT INTO users(id, name, mail, imagePath, password) VALUE(?,?,?,?,?)`

	_, err := infrastructure.Db.ExecContext(context.TODO(), sql, id, mail, imagePath, password)

	if err != nil {
		return fmt.Errorf("infra/create:%v", err)
	}

	return nil
}

func (a userRepository) Read() {

}
