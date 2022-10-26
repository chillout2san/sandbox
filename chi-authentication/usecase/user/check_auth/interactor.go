package checkauth

import "chi_sample/common/types"

func Interact() OutPut {
	return OutPut{
		IsAuthenticated: true,
		ErrType:         types.NO_ERROR,
		ErrMessage:      "",
	}
}
