package config

import "os"

// 環境変数をパースした構造体
type env struct {
	AllowOrigin string
	DbUser      string
	DbPassword  string
	DbHost      string
	DbName      string
}

var Enviroment *env

func init() {
	Enviroment = &env{
		AllowOrigin: os.Getenv("ALLOW_ORIGIN"),
		DbUser:      os.Getenv("MYSQL_USER"),
		DbPassword:  os.Getenv("MYSQL_PASSWORD"),
		DbHost:      os.Getenv("MYSQL_HOST"),
		DbName:      os.Getenv("MYSQL_DB_NAME"),
	}
}
