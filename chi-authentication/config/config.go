package config

import "os"

type Env struct {
	AllowOrigin string
}

var Enviroment *Env

func init() {
	ao := os.Getenv("ALLOW_ORIGIN")
	Enviroment = &Env{
		AllowOrigin: ao,
	}
}
