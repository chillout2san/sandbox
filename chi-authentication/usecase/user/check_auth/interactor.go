package checkauth

import "chi_sample/common/types"

func Interact() outPut {
	return outPut{
		IsAuthenticated: true,
		ErrType:         types.NO_ERROR,
		ErrMessage:      "",
	}
}
