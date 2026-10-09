package login

import (
	"chi_sample/common/utils"
	"chi_sample/domain/user"
	"context"
)

type loginUseCase struct {
	userRepository user.IUserRepository
}

// loginのinteractorを返却する
func NewLoginUseCase(ur user.IUserRepository) loginUseCase {
	return loginUseCase{userRepository: ur}
}

func (li loginUseCase) Execute(ctx context.Context, i InputDto) OutputDto {
	u, err := li.userRepository.GetByMail(ctx, i.Mail)
	if err != nil {
		return OutputDto{
			Id:         "",
			Token:      "",
			ErrMessage: err.Error(),
		}
	}

	if u.Id() == "" {
		return OutputDto{
			Id:         "",
			Token:      "",
			ErrMessage: "ユーザーが存在しません。",
		}
	}

	p, err := li.userRepository.GetPassByMail(ctx, i.Mail)
	if err != nil {
		return OutputDto{
			Id:         "",
			Token:      "",
			ErrMessage: err.Error(),
		}
	}

	if utils.CreateHash(i.Password) != p.Value {
		return OutputDto{
			Id:         "",
			Token:      "",
			ErrMessage: "パスワードが異なります。",
		}
	}

	token, err := utils.CreateJwt(u.Id())
	if err != nil {
		return OutputDto{
			Id:         "",
			Token:      "",
			ErrMessage: err.Error(),
		}
	}

	return OutputDto{
		Id:         u.Id(),
		Token:      token,
		ErrMessage: "",
	}
}
