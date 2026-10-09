package main

import (
	"chi_sample/infrastructure"
	"chi_sample/presentation"
	"log"
	"net/http"
)

func main() {
	router := presentation.NewRouter()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("main.http.ListenAndServe failed:%v;", err)
	}

	defer infrastructure.Db.Close()
}
