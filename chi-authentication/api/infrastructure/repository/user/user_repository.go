package user

import (
	duser "chi_sample/domain/user"
	"context"
	"database/sql"
	"fmt"
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
	tx, err := ur.db.Begin()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("userRepository.Create db.Begin failed:%v;", err)
	}

	sql := `INSERT INTO users(id, name, mail, imagePath, pass) VALUE(?,?,?,?,?)`

	if _, err = tx.ExecContext(ctx, sql, u.Id(), u.Name(), u.Mail(), u.ImagePath(), p.Value); err != nil {
		tx.Rollback()
		return fmt.Errorf("userRepository.Create ExecContext failed:%v;", err)
	}

	tx.Commit()
	return nil
}

// メールアドレスをキーとして、登録されているユーザー情報を取得する
func (ur userRepository) GetByMail(ctx context.Context, value string) (duser.User, error) {
	tx, err := ur.db.Begin()
	if err != nil {
		tx.Rollback()
		return duser.Reconstruct("", "", "", ""), fmt.Errorf("userRepository.GetByMail beginning of transaction failed:%v;", err)
	}

	sql := `SELECT id, name, mail, imagePath FROM users WHERE mail=?`

	row, err := ur.db.QueryContext(ctx, sql, value)
	if err != nil {
		tx.Rollback()
		return duser.Reconstruct("", "", "", ""), fmt.Errorf("userRepository.GetByMail QueryContext failed:%v;", err)
	}

	defer row.Close()

	var (
		id, name, mail, imagePath string
	)

	for row.Next() {
		if err := row.Scan(&id, &name, &mail, &imagePath); err != nil {
			tx.Rollback()
			return duser.Reconstruct("", "", "", ""), fmt.Errorf("userRepository.GetByMail rows.Scan failed:%v;", err)
		}
	}

	user := duser.Reconstruct(id, name, mail, imagePath)
	tx.Commit()
	return user, nil
}

// メールアドレスをキーとして、登録されているパスワードのハッシュ値を取得する
func (ur userRepository) GetPassByMail(ctx context.Context, value string) (duser.Password, error) {
	tx, err := ur.db.Begin()
	if err != nil {
		tx.Rollback()
		return duser.Password{}, fmt.Errorf("userRepository.GetPassByMail beginning of transaction failed:%v;", err)
	}

	sql := `SELECT pass FROM users WHERE mail=?`

	row, err := tx.QueryContext(ctx, sql, value)
	if err != nil {
		tx.Rollback()
		return duser.Password{}, fmt.Errorf("userRepository.GetPassByMail.row.Scan failed:%v;", err)
	}

	defer row.Close()

	var pass string

	for row.Next() {
		if err := row.Scan(&pass); err != nil {

			tx.Rollback()
			return duser.Password{}, fmt.Errorf("userRepository.GetPassByMail.row.Scan failed:%v;", err)
		}
	}

	p := duser.ReconstructPassWord(pass)
	tx.Commit()
	return p, nil
}
