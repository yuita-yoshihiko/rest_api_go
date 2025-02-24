package main

import (
	"net/http"
	"rest_api_go/infrastructure/router"
)

func main() {
	http.ListenAndServe(":8089", router.NewRouter())
}
