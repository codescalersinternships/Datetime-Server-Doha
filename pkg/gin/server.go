package pkg

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DateTime(c *gin.Context) {
	// returns if the method wasn't get
	if c.Request.Method != "GET" {
		c.Status(http.StatusMethodNotAllowed)
		return
	}

	// writes date by unix format
	fmt.Fprint(c.Writer, time.Now().Format(time.UnixDate))

	// returns status code 200
	c.Status(http.StatusOK)
}
