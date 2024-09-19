package server

import (
	"net/http"
	"net/http/httptest"

	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestServer(t *testing.T) {

	testcase := []struct {
		description string
		status      string
		path        string
		method      string
		expect      string
	}{
		{
			description: "valid request + matched time response",
			status:      "200 OK",
			path:        "/datetime",
			method:      "GET",
			expect:      time.Now().Format(time.UnixDate),
		},
		{
			description: "invalid method request",
			status:      "405 Method Not Allowed",
			path:        "/datatime",
			method:      "POST",
			expect:      "",
		},
	}

	for _, test := range testcase {
		t.Run(test.description, func(t *testing.T) {
			request, err := http.NewRequest(test.method, test.path, nil)
			if err != nil {
				t.Logf("%s", err.Error())
			}
			response := httptest.NewRecorder()

			DateTimeHandler(response, request)

			assert.Equal(t, test.expect, response.Body.String())
			assert.Equal(t, test.status, response.Result().Status)
		})
	}

}
