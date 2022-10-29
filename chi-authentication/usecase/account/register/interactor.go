package register

import (
	"chi_sample/domain/service"
	"chi_sample/domain/user"
)

type accountInteractor struct {
	userRepository user.IUserRepository
}

// accountのinteractorを返却する
func NewAccountInteractor(ui user.IUserRepository) accountInteractor {
	return accountInteractor{userRepository: ui}
}

func (ai accountInteractor) Interact(i InputDto) OutputDto {
	u, err := user.NewUser(i.Name, i.Mail, i.ImagePath)

	if err != nil {
		return OutputDto{
			IsRegistered: false,
			ErrMessage:   err.Error(),
		}
	}

	p, err := user.NewPassword(i.Password)

	if err != nil {
		return OutputDto{
			IsRegistered: false,
			ErrMessage:   err.Error(),
		}
	}

	isRegistered, err := service.CheckRegistered(ai.userRepository, u.Mail)

	if err != nil {
		return OutputDto{
			IsRegistered: false,
			ErrMessage:   err.Error(),
		}
	}

	if isRegistered {
		return OutputDto{
			IsRegistered: false,
			ErrMessage:   "既に会員登録されているメールアドレスです。",
		}
	}

	err = ai.userRepository.Create(u, p)

	if err != nil {
		return OutputDto{
			IsRegistered: false,
			ErrMessage:   err.Error(),
		}
	}

	return OutputDto{
		IsRegistered: true,
		ErrMessage:   "",
	}
}
