package user

import (
	"context"

	"chi_sample/domain/user"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (mur MockUserRepository) Create(ctx context.Context, u user.User, p user.Password) error {
	args := mur.Called(ctx, u, p)
	return args.Error(0)
}

func (mur MockUserRepository) GetByMail(ctx context.Context, value string) (user.User, error) {
	args := mur.Called(ctx, value)
	return args.Get(0).(user.User), args.Error(1)
}

func (mur MockUserRepository) GetPassByMail(ctx context.Context, value string) (user.Password, error) {
	args := mur.Called(ctx, value)
	return args.Get(0).(user.Password), args.Error(1)
}
