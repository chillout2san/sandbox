package login

type InputDto struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type OutputDto struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ImagePath  string `json:"imagePath"`
	Token      string `json:"token"`
	ErrMessage string `json:"errMessage"`
}
