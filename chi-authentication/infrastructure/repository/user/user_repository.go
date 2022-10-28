package user

import (
	"chi_sample/infrastructure"
	"context"
	"errors"
	"log"
)

type userRepository struct {
}

// ユーザーリポジトリを返却する
func NewUserRepository() userRepository {
	return userRepository{}
}

// 新しくユーザーを登録する。
func (u userRepository) Create(id, name, mail, imagePath, password string) error {
	sql := `INSERT INTO users(id, name, mail, imagePath, pass) VALUE(?,?,?,?,?)`

	_, err := infrastructure.Db.ExecContext(context.TODO(), sql, id, name, mail, imagePath, password)

	if err != nil {
		log.Println("userRepository.Create failed:", err)
		return errors.New("ユーザー登録できませんでした。")
	}

	return nil
}
