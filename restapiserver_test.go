// run test: go test -v
package goherokuauth

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthPostHandler(t *testing.T) {

	r := httptest.NewRequest("POST", "/api/v1/auth", strings.NewReader("{\"Username\": \"Me\", \"PasswordHash\": \"Me\"}"))
	w := httptest.NewRecorder()
	h := http.HandlerFunc(authRestAPIHandler)

	h.ServeHTTP(w, r)

	// check HTTP response status code
	wantCode := http.StatusInternalServerError
	if s := w.Code; s != wantCode {
		t.Errorf("handler return wrong status code: got %v, want %v", s, wantCode)
	}

	// check HTTP response header content type
	wantContentType := "text/plain; charset=utf-8"
	if c := w.Header().Get("Content-type"); c != wantContentType {
		t.Errorf("handler return wrong status code: got %v, want %v", c, wantContentType)
	}

	// check HTTP response body
	wantBody := "Internal Server Error\n"
	if w.Body.String() != wantBody {
		t.Errorf("handler return wrong status code: got %v, want %v", w.Body.String(), wantBody)
	}
}

func TestStatusNotFoundHandler(t *testing.T) {

	r := httptest.NewRequest("DELETE", "/api/v1/helloworld", nil)
	w := httptest.NewRecorder()
	h := http.HandlerFunc(authRestAPIHandler)

	h.ServeHTTP(w, r)

	// check HTTP response status code
	wantCode := http.StatusNotFound
	if s := w.Code; s != wantCode {
		t.Errorf("handler return wrong status code: got %v, want %v", s, wantCode)
	}

	// check HTTP response header content type
	wantContentType := "text/plain; charset=utf-8"
	if c := w.Header().Get("Content-type"); c != wantContentType {
		t.Errorf("handler return wrong status code: got %v, want %v", c, wantContentType)
	}

	// check HTTP response body
	wantBody := "Your requested method (DELETE) is not found"
	if w.Body.String() != wantBody {
		t.Errorf("handler return wrong status code: got %v, want %v", w.Body.String(), wantBody)
	}
}
