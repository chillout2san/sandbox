package checkauth

import "chi_sample/common/types"

type OutPut struct {
	IsAuthenticated bool
	ErrType         types.ErrType
	ErrMessage      string
}
