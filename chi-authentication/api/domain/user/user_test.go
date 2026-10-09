package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	testcases := []struct {
		label     string
		id        string
		userName  string
		mail      string
		imagePath string
		want1     User
		want2     error
	}{
		{
			label:     "失敗:idが空",
			id:        "",
			userName:  "hogeName",
			mail:      "hoge@example.com",
			imagePath: "https://hoge.com",
			want1:     &user{},
			want2:     errors.New("idが空です。"),
		},
		{
			label:     "失敗:名前が空",
			id:        "testid",
			userName:  "",
			mail:      "hoge@example.com",
			imagePath: "https://hoge.com",
			want1:     &user{},
			want2:     errors.New("名前が空です。"),
		},
		{
			label:     "失敗:メールアドレスが空",
			id:        "testid",
			userName:  "hogeName",
			mail:      "",
			imagePath: "https://hoge.com",
			want1:     &user{},
			want2:     errors.New("メールアドレスが空です。"),
		},
		{
			label:     "成功",
			id:        "testid",
			userName:  "hoge",
			mail:      "hoge@example.com",
			imagePath: "https://hoge.com",
			want1: &user{
				id:        "testid",
				name:      "hoge",
				mail:      "hoge@example.com",
				imagePath: "https://hoge.com",
			},
			want2: nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.label, func(t *testing.T) {
			got, err := New(tt.id, tt.userName, tt.mail, tt.imagePath)

			assert.Equal(t, tt.want1, got)
			assert.Equal(t, tt.want2, err)
		})
	}
}

func TestMappedUser(t *testing.T) {
	testcases := []struct {
		label     string
		id        string
		userName  string
		mail      string
		imagePath string
		want1     User
	}{
		{
			label:     "成功",
			id:        "testid",
			userName:  "hoge",
			mail:      "hoge@example.com",
			imagePath: "https://hoge.com",
			want1: &user{
				id:        "testid",
				name:      "hoge",
				mail:      "hoge@example.com",
				imagePath: "https://hoge.com",
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.label, func(t *testing.T) {
			got := Reconstruct(tt.id, tt.userName, tt.mail, tt.imagePath)

			assert.Equal(t, tt.want1, got)
		})
	}
}
