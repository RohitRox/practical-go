package controllers

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestByeController(t *testing.T) {
	t.Run("Get /bye", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/bye", nil)

		rec := httptest.NewRecorder()
		handler := http.HandlerFunc(ByeController)

		handler.ServeHTTP(rec, req)

		expected, got := "Bye!", rec.Body.String()

		if expected != got {
			t.Errorf("Expected: %s Got: %s", expected, got)
		}
	})
}
