package user

import (
	"chi_sample/common/utils"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPassword(t *testing.T) {
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
