package checkauth

import "chi_sample/common/types"

type outPut struct {
	IsAuthenticated bool
	ErrType         types.ErrType
	ErrMessage      string
}
