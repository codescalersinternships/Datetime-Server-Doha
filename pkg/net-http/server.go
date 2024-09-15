package pkg

import (
	"fmt"
	"net/http"
	"time"
)


func DateTime(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return 
	}

	fmt.Fprint(w, time.Now().Format(time.UnixDate))

	w.WriteHeader(http.StatusOK)
}
