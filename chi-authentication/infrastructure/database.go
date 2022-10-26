package infrastructure

import (
	"chi_sample/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	e := config.Enviroment
	dsn := fmt.Sprintf("%s:%s@cloudsql(%s)/%s", e.DbUser, e.DbPassword, e.DbHost, e.DbName)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Printf("データベースとの接続に失敗しました。:%v\n", err)
		return
	}

	Db = db
	log.Println("データベースとの接続に成功しました。")
}
