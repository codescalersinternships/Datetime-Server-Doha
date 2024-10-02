package server

import (
	"net/http"
	"net/http/httptest"
	"reflect"

	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func assertIsEqual(t *testing.T, type1, type2 any) {
	t.Helper()

	if !reflect.DeepEqual(type1, type2) {
		t.Errorf("i expect %v, found %v", type1, type2)
	}
}

func TestServer(t *testing.T) {

	testcase := []struct {
		description string
		status      int
		path        string
		method      string
		expect      string
	}{
		{
			description: "valid request",
			status:      200,
			path:        "/datetime",
			method:      "GET",
			expect:      time.Now().Format(time.UnixDate),
		},
		{
			description: "invalid method request",
			status:      404,
			path:        "/datetime",
			method:      "POST",
			expect:      "404 page not found",
		},
	}

	for _, test := range testcase {
		t.Run(test.description, func(t *testing.T) {
			r := SetUpRouter()
			r.GET(test.path, DateTimeHandler)
			req, _ := http.NewRequest(test.method, test.path, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assertIsEqual(t, test.expect, w.Body.String())
			assertIsEqual(t, test.status, w.Code)
		})
	}

}
