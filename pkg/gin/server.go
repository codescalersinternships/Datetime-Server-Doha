package pkg

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func DateTime(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.Status(http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(c.Writer, time.Now().Format(time.UnixDate))

	c.Status(http.StatusOK)
}
