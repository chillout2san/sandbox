package infrastructure

import (
	"chi_sample/config"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	if config.Environment.DbFlag == "GCP" {
		setUpCloudSQL()
		return
	}
	setUpDockerMySQL()
}

func setUpCloudSQL() {
	log.Println("Trying db connection by setUpCloudSQL...")
	e := config.Environment

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		log.Fatalf("CloudSQLConn NewDaialer failed:%v;", err)
	}

	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, e.DbHost)
		})

	dsn := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true",
		e.DbUser, e.DbPass, e.DbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Database connection failed:%v;", err)
	}

	Db = db
	log.Println("Database connection success")
}

func setUpDockerMySQL() {
	fmt.Println("Trying db connection by setUpDockerMySQL...")
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println("time.LoadLocation failed:", err)
		return
	}

	// dsnを構造体で構築
	c := mysql.Config{
		DBName:    config.Environment.DbName,
		User:      config.Environment.DbUser,
		Passwd:    config.Environment.DbPass,
		Addr:      config.Environment.DbHost,
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	// FormatDSNメソッドで文字列に変換
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatalf("Database connection failed:%v;", err)
	}

	Db = db
}
