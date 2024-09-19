package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// DateTime writes the current time in response with its status code
func DateTimeHandler(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.Status(http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(c.Writer, time.Now().Format(time.UnixDate))

	c.Status(http.StatusOK)
}
