package login

type InputDto struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type OutputDto struct {
	Id         string `json:"id"`
	Token      string `json:"token"`
	ErrMessage string `json:"errMessage"`
}
