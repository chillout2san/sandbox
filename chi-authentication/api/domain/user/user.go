package user

import (
	"errors"
)

// user構造体をexportしたくないが、型情報として何かしらexportしたい
// そのためgetterを持ったinterfaceをUserモデルの型とする
type User interface {
	Id() string
	Name() string
	Mail() string
	ImagePath() string
	SetName(name string)
	SetMail(mail string)
	SetImagePath(imagePath string)
}

// ユーザーモデル
type user struct {
	id        string // 一意のid
	name      string // ユーザーの名前
	mail      string // ユーザーのメールアドレス
	imagePath string // ユーザーの画像のパス
}

func (u *user) Id() string {
	return u.id
}

func (u *user) Name() string {
	return u.name
}

func (u *user) Mail() string {
	return u.mail
}

func (u *user) ImagePath() string {
	return u.imagePath
}

func (u *user) SetName(name string) {
	u.name = name
}

func (u *user) SetMail(mail string) {
	u.mail = mail
}

func (u *user) SetImagePath(imagePath string) {
	u.imagePath = imagePath
}

// 新規ユーザーを作成時に用いるファクトリ関数
func New(id string, name string, mail string, imagePath string) (User, error) {
	if id == "" {
		return &user{}, errors.New("idが空です。")
	}

	if name == "" {
		return &user{}, errors.New("名前が空です。")
	}

	if mail == "" {
		return &user{}, errors.New("メールアドレスが空です。")
	}

	// TODO: imagePathがURLとして正しい形か判断したい

	return &user{
		id:        id,
		name:      name,
		mail:      mail,
		imagePath: imagePath,
	}, nil
}

// 既存ユーザーのパース時に用いるファクトリ関数
func Reconstruct(id string, name string, mail string, imagePath string) User {
	return &user{
		id:        id,
		name:      name,
		mail:      mail,
		imagePath: imagePath,
	}
}
