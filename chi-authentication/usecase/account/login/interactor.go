package login

import (
	"chi_sample/common/utils"
	"chi_sample/domain/user"
)

type loginInteractor struct {
	userRepository user.IUserRepository
}

// loginのinteractorを返却する
func NewLoginInteractor(ur user.IUserRepository) loginInteractor {
	return loginInteractor{userRepository: ur}
}

func (li loginInteractor) Interact(i InputDto) OutputDto {
	u, err := li.userRepository.GetByMail(i.Mail)

	if err != nil {
		return OutputDto{
			Id:         "",
			Name:       "",
			ImagePath:  "",
			Token:      "",
			ErrMessage: err.Error(),
		}
	}

	p, err := li.userRepository.GetPassByMail(i.Mail)

	if err != nil {
		return OutputDto{
			Id:         "",
			Name:       "",
			ImagePath:  "",
			Token:      "",
			ErrMessage: err.Error(),
		}
	}

	if utils.CreateHash(i.Password) != p.Value {
		return OutputDto{
			Id:         "",
			Name:       "",
			ImagePath:  "",
			Token:      "",
			ErrMessage: "パスワードが異なります。",
		}
	}

	return OutputDto{
		Id:         u.Id,
		Name:       u.Name,
		ImagePath:  u.ImagePath,
		Token:      "正しいtoken",
		ErrMessage: "",
	}
}
