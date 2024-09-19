package server

import (
	"net/http"
	"net/http/httptest"
	"reflect"

	"testing"
	"time"

)


func assertIsEqual(t *testing.T, type1, type2 any) {
	t.Helper() 

	if !reflect.DeepEqual(type1, type2) {
		t.Errorf("i expect %v, found %v", type1, type2)
	}
}

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

			assertIsEqual(t, test.expect, response.Body.String())
			assertIsEqual(t, test.status, response.Result().Status)
		})
	}

}
