package register

type InputDto struct {
	Name      string `json:"name"`
	Mail      string `json:"mail"`
	ImagePath string `json:"imagePath"`
	Password  string `json:"password"`
}

type OutputDto struct {
	IsRegistered bool   `json:"isRegistered"`
	ErrMessage   string `json:"errMessage"`
}
