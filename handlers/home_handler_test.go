package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {

	t.Run("Home page returns Hello World!", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(HomeHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

	})
}
