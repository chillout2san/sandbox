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

func TestExecute2(t *testing.T) {
	t.Run("失敗:GetByMailがエラーの場合", func(t *testing.T) {
		m := new(user.MockUserRepository)
		m.On("GetByMail", mock.Anything, mock.Anything).Return(dUser.Reconstruct("", "", "", ""), errors.New("GetByMailがエラー"))
		usecase := NewLoginUseCase(m)

		result := usecase.Execute(context.Background(), InputDto{})

		assert.Equal(t, result, OutputDto{ErrMessage: "GetByMailがエラー"})
	})
}
