package infrastructure

import (
	"chi_sample/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	e := config.Enviroment
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", e.DbUser, e.DbPassword, e.DbHost, e.DbPort, e.DbName)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("Fail to Open Db:%v\n", err)
		return
	}

	Db = db
}
