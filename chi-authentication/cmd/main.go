package main

import (
	"chi_sample/presentation"
	"net/http"
)

func main() {
	router := presentation.CreateRoute()

	http.ListenAndServe(":8080", router)
}
