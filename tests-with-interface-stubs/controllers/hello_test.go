package controllers

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestHelloController(t *testing.T) {
	t.Run("Get /hello", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/hello", nil)

		rec := httptest.NewRecorder()
		handler := http.HandlerFunc(HelloController)

		handler.ServeHTTP(rec, req)

		expected, got := "Hello!", rec.Body.String()

		if expected != got {
			t.Errorf("Expected: %s Got: %s", expected, got)
		}
	})
}
