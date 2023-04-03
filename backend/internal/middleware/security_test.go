package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var calledHandler bool

func TestSecurityMiddleware(t *testing.T) {
	// create a dummy handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// mark that the handler was called
		calledHandler = true
	})

	// create a new request and recorder
	req, err := http.NewRequest("GET", "https://api.fadli.dev/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// create a new SecurityMiddleware and wrap the dummy handler
	middleware := SecurityMiddleware(handler)

	// call the middleware with the request and recorder
	middleware.ServeHTTP(rr, req)

	// check the content type header
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Content-Type header = %q; want %q", ct, "application/json")
	}

	// check the handler was called
	if !calledHandler {
		t.Error("Handler was not called")
	}
}
