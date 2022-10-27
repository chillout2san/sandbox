package config

import "os"

// 環境変数をパースした構造体
type env struct {
	AllowOrigin string // corsで許可するオリジン
	DbUser      string // GCPのCloudSQLにアクセスするユーザー名
	DbPass      string // 該当ユーザーのパスワード
	DbHost      string // GCPのCloudSQLの接続名
	DbName      string // データベースの名前
}

var Enviroment *env

func init() {
	Enviroment = &env{
		AllowOrigin: os.Getenv("ALLOW_ORIGIN"),
		DbUser:      os.Getenv("DB_USER"),
		DbPass:      os.Getenv("DB_PASS"),
		DbHost:      os.Getenv("INSTANCE_CONNECTION_NAME"),
		DbName:      os.Getenv("DB_NAME"),
	}
}
