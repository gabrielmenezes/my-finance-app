package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielmenezes/financial-control/handlers"
)

func TestHealthCheck(t *testing.T){
  rr := httptest.NewRecorder()
  req := httptest.NewRequest(http.MethodGet, "/health", nil)

  handlers.HealthCheck(rr, req)

  if http.StatusOK != rr.Code {
    t.Errorf("status code expected: %d, got: %d", http.StatusOK, rr.Code)
  }
}
