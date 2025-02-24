package main

import (
	"net/http"
	"rest_api_go/infrastructure/db"
	"rest_api_go/infrastructure/router"
)

func main() {
	db.Init()
	http.ListenAndServe(":8089", router.NewRouter())
}
