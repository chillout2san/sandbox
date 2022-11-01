package user

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
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
			want1:     User{},
			want2:     errors.New("idが空です。"),
		},
		{
			label:     "失敗：名前が空",
			id:        "testid",
			userName:  "",
			mail:      "hoge@example.com",
			imagePath: "https://hoge.com",
			want1:     User{},
			want2:     errors.New("名前が空です。"),
		},
		{
			label:     "失敗：メールアドレスが空",
			id:        "testid",
			userName:  "hogeName",
			mail:      "",
			imagePath: "https://hoge.com",
			want1:     User{},
			want2:     errors.New("メールアドレスが空です。"),
		},
		{
			label:     "成功",
			id:        "testid",
			userName:  "hoge",
			mail:      "hoge@example.com",
			imagePath: "https://hoge.com",
			want1: User{
				Id:        "testid",
				Name:      "hoge",
				Mail:      "hoge@example.com",
				ImagePath: "https://hoge.com",
			},
			want2: nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.label, func(t *testing.T) {
			got, err := NewUser(tt.id, tt.userName, tt.mail, tt.imagePath)

			diff1 := cmp.Diff(tt.want1, got)
			var diff2 string

			if diff1 != "" {
				t.Errorf("%s failed, diff1: %v", tt.label, diff1)
			}

			if err != nil && tt.want2 != nil {
				diff2 = cmp.Diff(err.Error(), tt.want2.Error())
			}

			if diff2 != "" {
				t.Errorf("%s failed, diff2: %v", tt.label, diff2)
			}
		})
	}
}
