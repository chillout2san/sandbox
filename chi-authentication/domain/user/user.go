package user

import (
	"chi_sample/common/utils"
	"errors"
)

// ユーザーモデル
type User struct {
	Id        string // 一意のid
	Name      string // ユーザーの名前
	Mail      string // ユーザーのメールアドレス
	ImagePath string // ユーザーの画像のパス
}

// ユーザーモデルが正しくインスタンス化されたかどうかを返却する
func (u User) IsValid() bool {
	return u.Id != ""
}

// 新規ユーザーを作成時に用いるファクトリ関数
func NewUser(name string, mail string, imagePath string) (User, error) {
	if name == "" {
		return User{}, errors.New("名前が空です。")
	}

	if mail == "" {
		return User{}, errors.New("メールアドレスが空です。")
	}

	// TODO: imagePathがURLとして正しい形か判断したい

	id, err := utils.CreateUlid()

	if err != nil {
		return User{}, err
	}

	return User{
		Id:        id.String(),
		Name:      name,
		Mail:      mail,
		ImagePath: imagePath,
	}, nil
}

// 既存ユーザーのパース時に用いるファクトリ関数
func MappedUser(id string, name string, mail string, imagePath string) User {
	return User{
		Id:        id,
		Name:      name,
		Mail:      mail,
		ImagePath: imagePath,
	}
}
