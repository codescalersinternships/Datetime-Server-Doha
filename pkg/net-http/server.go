package pkg

import (
	"fmt"
	"net/http"
	"time"
)

func DateTimeHandler(w http.ResponseWriter, r *http.Request) {
	// returns if the method wasn't get
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// writes date by unix format
	fmt.Fprint(w, time.Now().Format(time.UnixDate))

	// returns status code 200
	w.WriteHeader(http.StatusOK)
}
