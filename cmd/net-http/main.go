package main

import (
	"net/http"

	pkghttp "github.com/dohaelsawy/codescalers/datetimeserver/pkg/net-http"
	pkgsh "github.com/dohaelsawy/codescalers/datetimeserver/pkg/shutdown"
)


func main() {

    server := http.Server{
        Addr: ":8090",
		Handler: http.HandlerFunc(pkghttp.DateTime),
    }

	pkgsh.ShutDown(&server)

}
