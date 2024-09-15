package pkg

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func DateTime(c *gin.Context){
	fmt.Fprintln(c.Writer,time.Now().Format("Jan 02, 2006"))
	
	fmt.Fprintln(c.Writer,time.Now().Format("3:04 PM"))
	
	c.Status(http.StatusOK)
}