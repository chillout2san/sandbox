package user

import "context"

type IUserRepository interface {
	Create(ctx context.Context, u User, p Password) error
	GetByMail(ctx context.Context, value string) (User, error)
	GetPassByMail(ctx context.Context, value string) (Password, error)
}
