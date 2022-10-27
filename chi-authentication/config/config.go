package config

import "os"

// 環境変数をパースした構造体
type env struct {
	AllowOrigin string // corsで許可するオリジン
	DbUser      string // GCPのCloudSQLにアクセスするユーザー名
	DbPassword  string // 該当ユーザーのパスワード
	DbHost      string // GCPのCloudSQLの接続名
	DbName      string // データベースの名前
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
