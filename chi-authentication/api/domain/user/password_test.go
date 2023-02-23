package user

import (
	"chi_sample/common/utils"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	testcases := []struct {
		label string
		pass  string
		want1 Password
		want2 error
	}{
		{
			label: "失敗:passが空",
			pass:  "",
			want1: Password{},
			want2: errors.New("パスワードが空です。"),
		},
		{
			label: "成功",
			pass:  "hogepass",
			want1: Password{
				Value: utils.CreateHash("hogepass"),
			},
			want2: nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.label, func(t *testing.T) {
			got, err := NewPassword(tt.pass)

			assert.Equal(t, tt.want1, got)
			assert.Equal(t, tt.want2, err)
		})
	}
}

func TestMappedPassword(t *testing.T) {
	testcases := []struct {
		label string
		pass  string
		want1 Password
	}{
		{
			label: "成功",
			pass:  "hogepass",
			want1: Password{
				Value: "hogepass",
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.label, func(t *testing.T) {
			got := ReconstructPassWord(tt.pass)

			assert.Equal(t, tt.want1, got)
		})
	}
}
