package user

import (
	duser "chi_sample/domain/user"
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
func (ur userRepository) Create(ctx context.Context, u duser.User, p duser.Password) error {
	sql := `INSERT INTO users(id, name, mail, imagePath, pass) VALUE(?,?,?,?,?)`

	_, err := ur.db.ExecContext(ctx, sql, u.Id(), u.Name(), u.Mail(), u.ImagePath(), p.Value)

	if err != nil {
		log.Println("userRepository.Create failed:", err)
		return errors.New("ユーザー登録できませんでした。")
	}

	return nil
}

// メールアドレスをキーとして、登録されているユーザー情報を取得する
func (ur userRepository) GetByMail(ctx context.Context, value string) (duser.User, error) {
	sql := `SELECT id, name, mail, imagePath FROM users WHERE mail=?`

	row, err := ur.db.QueryContext(ctx, sql, value)

	if err != nil {
		log.Println("userRepository.GetByMail.QueryContext failed:", err)
		return duser.Reconstruct("", "", "", ""), errors.New("ユーザー情報を取得できませんでした。")
	}

	defer row.Close()

	var (
		id, name, mail, imagePath string
	)

	for row.Next() {
		if err := row.Scan(&id, &name, &mail, &imagePath); err != nil {
			log.Println("userRepository.GetByMail.rows.Scan failed:", err)
			return duser.Reconstruct("", "", "", ""), errors.New("ユーザー情報を取得できませんでした。")
		}
	}

	user := duser.Reconstruct(id, name, mail, imagePath)

	return user, nil
}

// メールアドレスをキーとして、登録されているパスワードのハッシュ値を取得する
func (ur userRepository) GetPassByMail(ctx context.Context, value string) (duser.Password, error) {
	sql := `SELECT pass FROM users WHERE mail=?`

	row, err := ur.db.QueryContext(ctx, sql, value)

	if err != nil {
		log.Println("userRepository.GetPassByMail.row.Scan failed", err)
		return duser.Password{}, errors.New("パスワード情報を取得できませんでした。")
	}

	defer row.Close()

	var pass string

	for row.Next() {
		if err := row.Scan(&pass); err != nil {
			log.Println("userRepository.GetPassByMail.row.Scan failed", err)
			return duser.Password{}, errors.New("パスワード情報を取得できませんでした。")
		}
	}

	p := duser.MappedPassword(pass)

	return p, nil
}
