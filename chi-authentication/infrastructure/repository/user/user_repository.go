package user

import (
	"chi_sample/domain/user"
	"context"
	"database/sql"
	"errors"
	"log"
)

type userRepository struct {
	db *sql.DB
}

// ユーザーリポジトリを返却する
func NewUserRepository(db *sql.DB) userRepository {
	return userRepository{db: db}
}

// 新しくユーザーを登録する
func (ur userRepository) Create(u user.User, p user.Password) error {
	sql := `INSERT INTO users(id, name, mail, imagePath, pass) VALUE(?,?,?,?,?)`

	_, err := ur.db.ExecContext(context.TODO(), sql, u.Id, u.Name, u.Mail, u.ImagePath, p.Value)

	if err != nil {
		log.Println("userRepository.Create failed:", err)
		return errors.New("ユーザー登録できませんでした。")
	}

	return nil
}

// メールアドレスをキーとして、登録されているユーザー情報を取得する
func (ur userRepository) GetByMail(value string) (user.User, error) {
	sql := `SELECT id, name, mail, imagePath FROM users WHERE mail=?`

	row := ur.db.QueryRowContext(context.TODO(), sql, value)

	var (
		id, name, mail, imagePath string
	)

	if err := row.Scan(&id, &name, &mail, &imagePath); err != nil {
		log.Println("userRepository.GetByMail.rows.Scan failed:", err)
		return user.User{}, errors.New("ユーザー情報を取得できませんでした。")
	}

	user := user.MappedUser(id, name, mail, imagePath)

	return user, nil
}

// メールアドレスをキーとして、登録されているパスワードのハッシュ値を取得する
func (ur userRepository) GetPassByMail(value string) (user.Password, error) {
	sql := `SELECT pass FROM users WHERE mail=?`

	row := ur.db.QueryRowContext(context.TODO(), sql, value)

	var pass string

	if err := row.Scan(&pass); err != nil {
		log.Println("userRepository.GetPassByMail.row.Scan failed", err)
		return user.Password{}, errors.New("パスワード情報を取得できませんでした。")
	}

	p := user.MappedPassword(pass)

	return p, nil
}
