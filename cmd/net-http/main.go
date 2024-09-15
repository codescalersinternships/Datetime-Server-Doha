package main

import (
	"log"
	"net/http"

	pkg "github.com/dohaelsawy/codescalers/datetimeserver/pkg/net-http"
)


func main() {

	http.HandleFunc("/datetime", pkg.DateTime)

	log.Println("Server started on port 8090")

	http.ListenAndServe(":8090", nil)
}
