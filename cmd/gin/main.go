package main

import (
	pkgin "github.com/dohaelsawy/codescalers/datetimeserver/pkg/gin"
	pkgsh "github.com/dohaelsawy/codescalers/datetimeserver/pkg/shutdown"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupHandler() *gin.Engine {
	router := gin.Default()
 
	router.GET("/datetime", pkgin.DateTime)

	return router
 }

func main() {

	server := http.Server{
		Addr: ":8080",
		Handler: setupHandler(),
	}

	pkgsh.ShutDown(&server)

}