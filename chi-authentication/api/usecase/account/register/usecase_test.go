package register

import (
	duser "chi_sample/domain/user"
	"chi_sample/infrastructure/repository/user"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExecute(t *testing.T) {
	t.Run("失敗:user.Newがエラーの場合", func(t *testing.T) {
		m := new(user.MockUserRepository)

		usecase := NewRegisterUseCase(m)
		result := usecase.Execute(context.Background(), InputDto{Name: ""})

		assert.Equal(t, result, OutputDto{IsRegistered: false, ErrMessage: "名前が空です。"})
	})

	t.Run("失敗:user.NewPasswordがエラーの場合", func(t *testing.T) {
		m := new(user.MockUserRepository)

		usecase := NewRegisterUseCase(m)
		result := usecase.Execute(context.Background(),
			InputDto{Name: "name", Mail: "mail", ImagePath: "imagePath", Password: ""})

		assert.Equal(t, result, OutputDto{IsRegistered: false, ErrMessage: "パスワードが空です。"})
	})

	t.Run("失敗:CheckRegisteredがエラーの場合", func(t *testing.T) {
		m := new(user.MockUserRepository)
		m.On("GetByMail", mock.Anything, mock.Anything).
			Return(duser.Reconstruct("", "", "", ""), errors.New("CheckRegisteredがエラー"))

		usecase := NewRegisterUseCase(m)
		result := usecase.Execute(context.Background(),
			InputDto{Name: "name", Mail: "mail", ImagePath: "imagePath", Password: "password"})

		assert.Equal(t, result, OutputDto{IsRegistered: false, ErrMessage: "CheckRegisteredがエラー"})
	})

	t.Run("失敗:既に会員登録されている場合", func(t *testing.T) {
		m := new(user.MockUserRepository)
		m.On("GetByMail", mock.Anything, mock.Anything).
			Return(duser.Reconstruct("id", "name", "mail", "imagePath"), nil)

		usecase := NewRegisterUseCase(m)
		result := usecase.Execute(context.Background(),
			InputDto{Name: "name", Mail: "mail", ImagePath: "imagePath", Password: "password"})

		assert.Equal(t, result, OutputDto{IsRegistered: false, ErrMessage: "既に会員登録されているメールアドレスです。"})
	})

	t.Run("失敗:userRepository.Createがエラーの場合", func(t *testing.T) {
		m := new(user.MockUserRepository)
		m.On("GetByMail", mock.Anything, mock.Anything).
			Return(duser.Reconstruct("", "", "", ""), nil)
		m.On("Create", mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("Createのエラー"))

		usecase := NewRegisterUseCase(m)
		result := usecase.Execute(context.Background(),
			InputDto{Name: "name", Mail: "mail", ImagePath: "imagePath", Password: "password"})

		assert.Equal(t, result, OutputDto{IsRegistered: false, ErrMessage: "Createのエラー"})
	})

	t.Run("成功", func(t *testing.T) {
		m := new(user.MockUserRepository)
		m.On("GetByMail", mock.Anything, mock.Anything).
			Return(duser.Reconstruct("", "", "", ""), nil)
		m.On("Create", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		usecase := NewRegisterUseCase(m)
		result := usecase.Execute(context.Background(),
			InputDto{Name: "name", Mail: "mail", ImagePath: "imagePath", Password: "password"})

		assert.Equal(t, result, OutputDto{IsRegistered: true, ErrMessage: ""})
	})
}
