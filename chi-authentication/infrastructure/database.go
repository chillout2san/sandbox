package infrastructure

import (
	"chi_sample/config"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	e := config.Enviroment

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		log.Println("CloudSQLConn NewDaialer failed:", err)
	}
	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, e.DbHost)
		})

	dsn := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true",
		e.DbUser, e.DbPass, e.DbName)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Println("Database connection failed:", err)
		return
	}

	Db = db
	log.Println("Database connection success")
}
