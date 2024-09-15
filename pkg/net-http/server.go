package pkg

import (
	"fmt"
	"net/http"
	"time"
)



func DateTime(w http.ResponseWriter , r *http.Request){

	fmt.Fprintln(w,time.Now().Format("Jan 02, 2006"))
	
	fmt.Fprintln(w,time.Now().Format("3:04 PM"))
	
	w.WriteHeader(http.StatusOK)
}