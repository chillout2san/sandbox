package login

import (
	"chi_sample/common/utils"
	dUser "chi_sample/domain/user"
	"chi_sample/infrastructure/repository/user"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExecute(t *testing.T) {
	t.Run("失敗:GetByMailがエラーの場合", func(t *testing.T) {
		m := new(user.MockUserRepository)
		m.On("GetByMail", mock.Anything, mock.Anything).
			Return(dUser.Reconstruct("", "", "", ""), errors.New("GetByMailがエラー"))

		usecase := NewLoginUseCase(m)
		result := usecase.Execute(context.Background(), InputDto{})

		assert.Equal(t, result, OutputDto{Id: "", Token: "", ErrMessage: "GetByMailがエラー"})
	})

	t.Run("失敗:GetByMailの戻り値のuserのidが空文字の場合", func(t *testing.T) {
		m := new(user.MockUserRepository)
		m.On("GetByMail", mock.Anything, mock.Anything).
			Return(dUser.Reconstruct("", "", "", ""), nil)

		usecase := NewLoginUseCase(m)
		result := usecase.Execute(context.Background(), InputDto{})

		assert.Equal(t, result, OutputDto{Id: "", Token: "", ErrMessage: "ユーザーが存在しません。"})
	})

	t.Run("失敗:パスワードのhash値が異なる場合", func(t *testing.T) {
		m := new(user.MockUserRepository)
		m.On("GetByMail", mock.Anything, mock.Anything).
			Return(dUser.Reconstruct("id", "name", "mail", "imagePath"), nil)
		m.On("GetPassByMail", mock.Anything, mock.Anything).
			Return(dUser.Password{Value: utils.CreateHash("correct")}, nil)

		usecase := NewLoginUseCase(m)
		result := usecase.Execute(context.Background(), InputDto{Password: "invalid"})

		assert.Equal(t, result, OutputDto{Id: "", Token: "", ErrMessage: "パスワードが異なります。"})
	})

	t.Run("成功", func(t *testing.T) {
		m := new(user.MockUserRepository)
		m.On("GetByMail", mock.Anything, mock.Anything).
			Return(dUser.Reconstruct("id", "name", "mail", "imagePath"), nil)
		m.On("GetPassByMail", mock.Anything, mock.Anything).
			Return(dUser.Password{Value: utils.CreateHash("correct")}, nil)

		usecase := NewLoginUseCase(m)
		result := usecase.Execute(context.Background(), InputDto{Password: "correct"})

		assert.Equal(t, result.Id, "id")
		assert.Equal(t, result.ErrMessage, "")
		assert.NotNil(t, result.Token)
	})
}
