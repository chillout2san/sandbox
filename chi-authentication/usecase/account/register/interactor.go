package register

import (
	"chi_sample/common/utils"
	"chi_sample/domain/service"
	"chi_sample/domain/user"
)

type registerInteractor struct {
	userRepository user.IUserRepository
}

// registerのinteractorを返却する
func NewRegisterInteractor(ui user.IUserRepository) registerInteractor {
	return registerInteractor{userRepository: ui}
}

func (ri registerInteractor) Interact(i InputDto) OutputDto {
	id, _ := utils.CreateUlid()
	u, err := user.NewUser(id.String(), i.Name, i.Mail, i.ImagePath)

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

	isRegistered := service.CheckRegistered(ri.userRepository, u.Mail)

	if isRegistered {
		return OutputDto{
			IsRegistered: false,
			ErrMessage:   "既に会員登録されているメールアドレスです。",
		}
	}

	err = ri.userRepository.Create(u, p)

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
