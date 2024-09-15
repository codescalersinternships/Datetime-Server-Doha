package main

import (
	pkg "github.com/dohaelsawy/codescalers/datetimeserver/pkg/gin"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/datetime",pkg.DateTime)

	r.Run("localhost:8080")

}
