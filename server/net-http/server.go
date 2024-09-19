package server

import (
	"fmt"
	"net/http"
	"time"
)


// DateTime writes in response with its status code
func DateTimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, time.Now().Format(time.UnixDate))

	w.WriteHeader(http.StatusOK)
}
