package login

import (
	dUser "chi_sample/domain/user"
	"chi_sample/infrastructure/repository/user"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExecute(t *testing.T) {
	type args struct {
		Id         string
		Name       string
		Mail       string
		ImagePath  string
		ErrMessage string
	}
	testcases := []struct {
		label string
		args  args
		want  OutputDto
	}{
		{
			label: "失敗:GetByMailがエラーの場合",
			args: args{
				Id:         "",
				Name:       "",
				Mail:       "",
				ImagePath:  "",
				ErrMessage: "GetByMailがエラー",
			},
			want: OutputDto{
				Id:         "",
				Token:      "",
				ErrMessage: "GetByMailがエラー",
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.label, func(t *testing.T) {
			m := new(user.MockUserRepository)
			m.On("GetByMail", mock.Anything, mock.Anything).Return(dUser.Reconstruct(tt.args.Id, tt.args.Name, tt.args.Mail, tt.args.ImagePath), errors.New(tt.args.ErrMessage))
			usecase := NewLoginUseCase(m)

			result := usecase.Execute(context.Background(), InputDto{})

			assert.Equal(t, result, tt.want)
		})
	}

}
