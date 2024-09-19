package main

import (
	"log"
	"net/http"

	serverGin "github.com/codescalersinternships/Datetime-Server-Doha/server/gin"
	serverShuts "github.com/codescalersinternships/Datetime-Server-Doha/server/shutdown"
	"github.com/gin-gonic/gin"
)

// setupHandler setup gin engin and call datetime handeler
func setupHandler() *gin.Engine {
	router := gin.Default()

	router.GET("/datetime", serverGin.DateTimeHandler)

	return router
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: setupHandler(),
	}

	go func() {
		log.Printf("Server listening on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	serverShuts.ShutDown(&server)

}
