package user

import (
	"chi_sample/common/utils"
	"errors"
)

type user struct {
	Id        string // 一意のid
	Name      string // ユーザーの名前
	Mail      string // ユーザーのメールアドレス
	ImagePath string // ユーザーの画像のパス
}

// 新規ユーザーを作成時に用いるファクトリ関数
func NewUser(name string, mail string, imagePath string) (user, error) {
	if name == "" {
		return user{}, errors.New("名前が空です。")
	}

	if mail == "" {
		return user{}, errors.New("メールアドレスが空です。")
	}

	// TODO: imagePathがURLとして正しい形か判断したい

	id, err := utils.CreateUlid()

	if err != nil {
		return user{}, err
	}

	return user{
		Id:        id.String(),
		Name:      name,
		Mail:      mail,
		ImagePath: imagePath,
	}, nil
}

// 既存ユーザーのパース時に用いるファクトリ関数
func MappedUser(id string, name string, mail string, imagePath string) (user, error) {
	return user{
		Id:        id,
		Name:      name,
		Mail:      mail,
		ImagePath: imagePath,
	}, nil
}
