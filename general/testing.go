package general

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// code based on https://go-chi.io/#/pages/testing

// ExecuteRequest creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
func ExecuteRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

// CheckResponseCode is a simple utility to check the response code
// of the response
func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
