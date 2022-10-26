package main

import (
	"chi_sample/infrastructure"
	"chi_sample/presentation"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := presentation.CreateRoute()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		e := fmt.Sprintf("APIサーバの起動に失敗しました。:%v\n", err)
		log.Println(e)
	}

	defer infrastructure.Db.Close()
}
